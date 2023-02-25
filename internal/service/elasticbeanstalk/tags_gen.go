// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package elasticbeanstalk

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// ListTags lists elasticbeanstalk service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn elasticbeanstalkiface.ElasticBeanstalkAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResourceWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.ResourceTags), nil
}

// []*SERVICE.Tag handling

// Tags returns elasticbeanstalk service tags.
func Tags(tags tftags.KeyValueTags) []*elasticbeanstalk.Tag {
	result := make([]*elasticbeanstalk.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &elasticbeanstalk.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from elasticbeanstalk service tags.
func KeyValueTags(ctx context.Context, tags []*elasticbeanstalk.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// UpdateTags updates elasticbeanstalk service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(ctx context.Context, conn elasticbeanstalkiface.ElasticBeanstalkAPI, identifier string, oldTagsMap interface{}, newTagsMap interface{}) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)
	removedTags := oldTags.Removed(newTags)
	updatedTags := oldTags.Updated(newTags)

	// Ensure we do not send empty requests
	if len(removedTags) == 0 && len(updatedTags) == 0 {
		return nil
	}

	input := &elasticbeanstalk.UpdateTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	if len(updatedTags) > 0 {
		input.TagsToAdd = Tags(updatedTags.IgnoreAWS())
	}

	if len(removedTags) > 0 {
		input.TagsToRemove = aws.StringSlice(removedTags.Keys())
	}

	_, err := conn.UpdateTagsForResourceWithContext(ctx, input)

	if err != nil {
		return fmt.Errorf("tagging resource (%s): %w", identifier, err)
	}

	return nil
}
