import React from 'react';
import Image from 'next/image';

interface QuickTransferCardProps {
  username: string;
  profilePicture: string;
}

export const QuickTransferCard: React.FC<QuickTransferCardProps> = ({ username, profilePicture }) => {
  return (
    <div className="flex items-center space-x-4 p-2 bg-white rounded-lg shadow">
      <Image
        src={profilePicture}
        alt={`${username}'s profile picture`}
        width={44}
        height={44}
        className="rounded-full object-cover"
      />
      <p className="font-semibold text-sm md:text-base">{username}</p>
    </div>
  );
};
