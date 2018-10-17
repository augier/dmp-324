'use strict';

// [START functions_helloworld_storage_generic]
/**
 * Generic background Cloud Function to be triggered by Cloud Storage.
 *
 * @param {object} event The Cloud Functions event.
 * @param {function} callback The callback function.
 */
exports.bucketListener = (event, callback) => {
    const file = event.data;
  
    console.log(`  File: ${file.name}`);
  
    callback();
  };
  // [END functions_helloworld_storage_generic]
  