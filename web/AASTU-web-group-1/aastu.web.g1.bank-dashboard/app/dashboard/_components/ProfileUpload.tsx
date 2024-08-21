"use client";

import { useState } from "react";
import { ref, uploadBytesResumable, getDownloadURL } from "firebase/storage";
import { storage } from "@/firebase";
import { Progress } from "@/components/ui/progress";
import { Button } from "@/components/ui/button";

const ProfileUpload = ({
  setProfilePictureUrl,
}: {
  setProfilePictureUrl: (url: string) => void;
}) => {
  const [file, setFile] = useState<File | null>(null);
  const [progress, setProgress] = useState(0);
  const [downloadURL, setDownloadURL] = useState<string | null>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const selectedFile = e.target.files[0];
      setFile(selectedFile);
      handleUpload(selectedFile); 
    }
  };

  const handleUpload = (file: File) => {
    const storageRef = ref(storage, `profiles/${file.name}`);
    const uploadTask = uploadBytesResumable(storageRef, file);

    uploadTask.on(
      "state_changed",
      (snapshot:any) => {
        const progress =
          (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
        setProgress(progress);
      },
      (error:any) => {
        console.error("Upload failed:", error);
      },
      () => {
        getDownloadURL(uploadTask.snapshot.ref).then((url:any) => {
          setDownloadURL(url);
          setProfilePictureUrl(url);
          console.log("File available at", url);
        });
      }
    );
  };

  return (
    <div className="flex flex-col  gap-4 p-4">
      <input
        accept="image/*"
        className="block border-dotted border-spacing-3 w-full text-sm text-gray-900 border border-gray-300 p-4 cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 rounded-2xl md:w-[400px]"
        id="file_input"
        type="file"
        onChange={handleFileChange}
      />
      {progress > 0 && <Progress value={progress} className="w-full" />}
      {downloadURL && (
        <div className="text-center mt-4">
          <p className="font-medium text-green-600">Upload complete!</p>
          <a
            href={downloadURL}
            target="_blank"
            rel="noopener noreferrer"
            className="text-blue-600 underline"
          >
            View Image
          </a>
        </div>
      )}
    </div>
  );
};

export default ProfileUpload;
