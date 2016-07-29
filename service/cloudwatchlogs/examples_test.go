// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatchlogs_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudWatchLogs_CancelExportTask() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.CancelExportTaskInput{
		TaskId: aws.String("ExportTaskId"), // Required
	}
	resp, err := svc.CancelExportTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_CreateExportTask() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.CreateExportTaskInput{
		Destination:         aws.String("ExportDestinationBucket"), // Required
		From:                aws.Int64(1),                          // Required
		LogGroupName:        aws.String("LogGroupName"),            // Required
		To:                  aws.Int64(1),                          // Required
		DestinationPrefix:   aws.String("ExportDestinationPrefix"),
		LogStreamNamePrefix: aws.String("LogStreamName"),
		TaskName:            aws.String("ExportTaskName"),
	}
	resp, err := svc.CreateExportTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_CreateLogGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.CreateLogGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_CreateLogStream() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
	}
	resp, err := svc.CreateLogStream(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DeleteDestination() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DeleteDestinationInput{
		DestinationName: aws.String("DestinationName"), // Required
	}
	resp, err := svc.DeleteDestination(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DeleteLogGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DeleteLogGroupInput{
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteLogGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DeleteLogStream() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DeleteLogStreamInput{
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
	}
	resp, err := svc.DeleteLogStream(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DeleteMetricFilter() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DeleteMetricFilterInput{
		FilterName:   aws.String("FilterName"),   // Required
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteMetricFilter(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DeleteRetentionPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DeleteRetentionPolicyInput{
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteRetentionPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DeleteSubscriptionFilter() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DeleteSubscriptionFilterInput{
		FilterName:   aws.String("FilterName"),   // Required
		LogGroupName: aws.String("LogGroupName"), // Required
	}
	resp, err := svc.DeleteSubscriptionFilter(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DescribeDestinations() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DescribeDestinationsInput{
		DestinationNamePrefix: aws.String("DestinationName"),
		Limit:     aws.Int64(1),
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.DescribeDestinations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DescribeExportTasks() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DescribeExportTasksInput{
		Limit:      aws.Int64(1),
		NextToken:  aws.String("NextToken"),
		StatusCode: aws.String("ExportTaskStatusCode"),
		TaskId:     aws.String("ExportTaskId"),
	}
	resp, err := svc.DescribeExportTasks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DescribeLogGroups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DescribeLogGroupsInput{
		Limit:              aws.Int64(1),
		LogGroupNamePrefix: aws.String("LogGroupName"),
		NextToken:          aws.String("NextToken"),
	}
	resp, err := svc.DescribeLogGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DescribeLogStreams() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        aws.String("LogGroupName"), // Required
		Descending:          aws.Bool(true),
		Limit:               aws.Int64(1),
		LogStreamNamePrefix: aws.String("LogStreamName"),
		NextToken:           aws.String("NextToken"),
		OrderBy:             aws.String("OrderBy"),
	}
	resp, err := svc.DescribeLogStreams(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DescribeMetricFilters() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DescribeMetricFiltersInput{
		LogGroupName:     aws.String("LogGroupName"), // Required
		FilterNamePrefix: aws.String("FilterName"),
		Limit:            aws.Int64(1),
		NextToken:        aws.String("NextToken"),
	}
	resp, err := svc.DescribeMetricFilters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_DescribeSubscriptionFilters() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.DescribeSubscriptionFiltersInput{
		LogGroupName:     aws.String("LogGroupName"), // Required
		FilterNamePrefix: aws.String("FilterName"),
		Limit:            aws.Int64(1),
		NextToken:        aws.String("NextToken"),
	}
	resp, err := svc.DescribeSubscriptionFilters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_FilterLogEvents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName:  aws.String("LogGroupName"), // Required
		EndTime:       aws.Int64(1),
		FilterPattern: aws.String("FilterPattern"),
		Interleaved:   aws.Bool(true),
		Limit:         aws.Int64(1),
		LogStreamNames: []*string{
			aws.String("LogStreamName"), // Required
			// More values...
		},
		NextToken: aws.String("NextToken"),
		StartTime: aws.Int64(1),
	}
	resp, err := svc.FilterLogEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_GetLogEvents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
		EndTime:       aws.Int64(1),
		Limit:         aws.Int64(1),
		NextToken:     aws.String("NextToken"),
		StartFromHead: aws.Bool(true),
		StartTime:     aws.Int64(1),
	}
	resp, err := svc.GetLogEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_PutDestination() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.PutDestinationInput{
		DestinationName: aws.String("DestinationName"), // Required
		RoleArn:         aws.String("RoleArn"),         // Required
		TargetArn:       aws.String("TargetArn"),       // Required
	}
	resp, err := svc.PutDestination(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_PutDestinationPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.PutDestinationPolicyInput{
		AccessPolicy:    aws.String("AccessPolicy"),    // Required
		DestinationName: aws.String("DestinationName"), // Required
	}
	resp, err := svc.PutDestinationPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_PutLogEvents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.PutLogEventsInput{
		LogEvents: []*cloudwatchlogs.InputLogEvent{ // Required
			{ // Required
				Message:   aws.String("EventMessage"), // Required
				Timestamp: aws.Int64(1),               // Required
			},
			// More values...
		},
		LogGroupName:  aws.String("LogGroupName"),  // Required
		LogStreamName: aws.String("LogStreamName"), // Required
		SequenceToken: aws.String("SequenceToken"),
	}
	resp, err := svc.PutLogEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_PutMetricFilter() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.PutMetricFilterInput{
		FilterName:    aws.String("FilterName"),    // Required
		FilterPattern: aws.String("FilterPattern"), // Required
		LogGroupName:  aws.String("LogGroupName"),  // Required
		MetricTransformations: []*cloudwatchlogs.MetricTransformation{ // Required
			{ // Required
				MetricName:      aws.String("MetricName"),      // Required
				MetricNamespace: aws.String("MetricNamespace"), // Required
				MetricValue:     aws.String("MetricValue"),     // Required
			},
			// More values...
		},
	}
	resp, err := svc.PutMetricFilter(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_PutRetentionPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.PutRetentionPolicyInput{
		LogGroupName:    aws.String("LogGroupName"), // Required
		RetentionInDays: aws.Int64(1),               // Required
	}
	resp, err := svc.PutRetentionPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_PutSubscriptionFilter() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.PutSubscriptionFilterInput{
		DestinationArn: aws.String("DestinationArn"), // Required
		FilterName:     aws.String("FilterName"),     // Required
		FilterPattern:  aws.String("FilterPattern"),  // Required
		LogGroupName:   aws.String("LogGroupName"),   // Required
		RoleArn:        aws.String("RoleArn"),
	}
	resp, err := svc.PutSubscriptionFilter(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchLogs_TestMetricFilter() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.TestMetricFilterInput{
		FilterPattern: aws.String("FilterPattern"), // Required
		LogEventMessages: []*string{ // Required
			aws.String("EventMessage"), // Required
			// More values...
		},
	}
	resp, err := svc.TestMetricFilter(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
