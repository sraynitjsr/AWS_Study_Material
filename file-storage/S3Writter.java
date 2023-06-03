import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.s3.AmazonS3;
import com.amazonaws.services.s3.AmazonS3ClientBuilder;
import com.amazonaws.services.s3.model.*;

import java.io.File;

public class S3Writter {
    public static void main(String[] args) {
        String accessKey = "sraynitjr";
        String secretKey = "tufan100";
        String bucketName = "my-bucket";
        String key = "my-object-key";
        String filePath = "$home/temp/file.txt";
        String destinationFilePath = "$home/temp/downloaded-file.txt";

        AWSCredentials credentials = new BasicAWSCredentials(accessKey, secretKey);

        AmazonS3 s3Client = AmazonS3ClientBuilder.standard()
                .withCredentials(new AWSStaticCredentialsProvider(credentials))
                .withRegion(Regions.us - east - 1)
                .build();

        // Create a bucket
        s3Client.createBucket(bucketName);

        // Upload an object to the bucket
        s3Client.putObject(bucketName, key, new File(filePath));

        // Download an object from the bucket
        s3Client.getObject(new GetObjectRequest(bucketName, key), new File(destinationFilePath));

        // Delete an object from the bucket
        s3Client.deleteObject(bucketName, key);

        // Delete the bucket
        s3Client.deleteBucket(bucketName);

        // List all objects in the bucket
        ListObjectsV2Request listObjectsRequest = new ListObjectsV2Request()
                .withBucketName(bucketName)
                .withMaxKeys(1000); // Maximum number of keys to retrieve in one request

        ListObjectsV2Result result;

        do {
            result = s3Client.listObjectsV2(listObjectsRequest);

            for (S3ObjectSummary objectSummary : result.getObjectSummaries()) {
                // Access the details of each object
                String objectKey = objectSummary.getKey();
                long size = objectSummary.getSize();
                System.out.println("Key: " + objectKey + ", Size: " + size);
            }

            // Set the next continuation token to retrieve the next set of keys
            String token = result.getNextContinuationToken();
            listObjectsRequest.setContinuationToken(token);
        } while (result.isTruncated());
    }
}
