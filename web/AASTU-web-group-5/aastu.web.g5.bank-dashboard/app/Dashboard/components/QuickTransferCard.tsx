'use client';

import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Image from 'next/image';
import { useSession } from 'next-auth/react';
import Profile from '@/public/images/Rectangle 65 (2).png'

interface QuickTransferCardProps {
  userId: string; 
  name: string;
  username: string;
  profilePicture: string;
  city: string;
  country: string;
}

export const QuickTransferCard: React.FC<QuickTransferCardProps> = ({ userId, name, username, profilePicture, city, country }) => {
  const { data: session } = useSession();
  const user = session?.user as { accessToken?: string; refreshToken?: string };

  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <div className="flex items-center space-x-4 p-2 rounded-lg ">
      {/* <Image
        src={Profile}
        alt={`${username}'s profile picture`}
        width={44}
        height={44}
        className="rounded-full object-cover"
      /> */}
      <div>
        <p className="font-semibold text-sm md:text-base">{name}</p>
        <p className="text-sm text-gray-500">@{username}</p>
        <p className="text-xs text-gray-400">{city}, {country}</p>
      </div>
    </div>
  );
};

export default QuickTransferCard;
