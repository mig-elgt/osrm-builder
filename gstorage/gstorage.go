package gstorage

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"cloud.google.com/go/storage"
	osrmbuilder "github.com/mig-elgt/osrm-builder"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type gstorage struct{}

func New() osrmbuilder.Uploader {
	return &gstorage{}
}

func (g *gstorage) Upload(bucket, root string) error {
	logrus.Infof("get file names from root dir builder: %v", root)
	files, err := g.getFileNames(root)
	if err != nil {
		return err
	}
	logrus.Infof("creating bucket: %v", bucket)
	if err := g.createBucket(bucket); err != nil {
		return err
	}
	// TODO: If there was an error do a rollback: delete bucket with objects and return an error
	if err := g.uploadAll(context.Background(), root, bucket, files); err != nil {
		return fmt.Errorf("could not upload OSM data to bucket: %v", bucket)
	}
	return nil
}

func (g *gstorage) createBucket(bucketName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "could not create new storage client")
	}
	defer client.Close()
	bkt := client.Bucket(bucketName)
	if err := bkt.Create(ctx, "ntrack-system", nil); err != nil {
		return errors.Wrap(err, "could not create bucket")
	}
	return nil
}

func (g *gstorage) getFileNames(root string) ([]string, error) {
	fileNames := []string{}
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read dir: %v", root)
	}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

func (g *gstorage) uploadAll(ctx context.Context, root, bucket string, files []string) error {
	errs, ctx := errgroup.WithContext(ctx)
	for _, fileName := range files {
		fileName := fileName
		errs.Go(func() error {
			if err := g.uploadFile(root, fileName, bucket); err != nil {
				logrus.Errorf("could not upload file %v: %v", fileName, err)
				return errors.Wrapf(err, "could not upload file %v", fileName)
			}
			logrus.Infof("file %v uploaded", fileName)
			return nil
		})
	}
	return errs.Wait()
}

func (g *gstorage) uploadFile(root, fileName, bucket string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "could not create storage client")
	}
	defer client.Close()

	// Open local file.
	f, err := os.Open(root + "/" + fileName)
	if err != nil {
		return errors.Wrapf(err, "could not open file: %v", fileName)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*10)
	defer cancel()
	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(fileName).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}
