import { storage } from '@/firebase'; // Import your Firebase storage instance
import { getDownloadURL, ref, uploadBytes } from 'firebase/storage';

export const uploadImage = async (file: File): Promise<string> => {
  const storageRef = ref(storage, `images/${file.name}`); // Create a reference to the file in storage

  try {
    const snapshot = await uploadBytes(storageRef, file); // Upload the file
    const downloadURL = await getDownloadURL(snapshot.ref); // Get the download URL
    return downloadURL;
  } catch (error) {
    console.error('Error uploading image:', error);
    throw error;
  }
};
