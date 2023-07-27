# Updating Services

In your API Gateway, file changes are tracked using the [fsnotify](https://github.com/fsnotify/fsnotify) library, which allows for efficient monitoring of changes in the file system.

## Adding new IDLs

To add a new Interface Definition Language (IDL) for your services, follow these steps:

1. **Write your new IDL file:** Begin by creating a new IDL file that defines the interface for your desired services.

2. **Transfer it into your IDL Directory:** Once you have written the IDL file, transfer it into the designated IDL Directory of your API Gateway. This directory serves as a repository for all IDL files.

3. **Validation of the IDL:** The API Gateway will automatically validate the new IDL file to ensure it adheres to the required syntax and structure. If the IDL is valid, the client will be updated with the new services, making them available for use.

4. **Handling invalid IDLs:** In case the new IDL file is found to be invalid during validation, you will need to review and rewrite it until it conforms to the correct IDL specifications. Once the issues are resolved, repeat the steps to add the updated IDL to your API Gateway.

## Modifying existing IDLs

If you need to make changes to an existing IDL to update or extend the services in your API Gateway, follow these steps:

1. **Open your existing IDL file:** Locate and open the IDL file that you want to modify using a text editor or an integrated development environment (IDE).

2. **Edit the IDL file:** Make the necessary changes to the IDL file, ensuring that the modifications adhere to the valid IDL syntax and do not introduce any errors.

3. **Validation of the modified IDL:** After editing the IDL file, the API Gateway will perform a validation check on the modified IDL. If the changes are valid, the client will be automatically updated with the new services reflecting the modifications.

4. **Handling invalid IDLs:** If the modified IDL contains errors or does not meet the required IDL specifications, you will need to correct the issues and revalidate the IDL. Repeat this process until the IDL file is valid and the services are updated accordingly.

By following these steps, you can effectively add new IDLs or modify existing ones to update the services offered by your API Gateway.

Last but not least, update the logic in `handler.go`.

### Service Registration and Discovery

Service registration and discovery is done using [etcd](https://etcd.io/docs/v3.5/)
and the [`registry-etcd`](https://github.com/kitex-contrib/registry-etcd) library.