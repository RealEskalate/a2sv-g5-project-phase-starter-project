'use client';

import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Image from 'next/image';

interface UserProfile {
  id: string;
  name: string;
  username: string;
  profilePicture: string;
}

interface QuickTransferCardProps {
  userId: string; // Pass user ID as a prop to fetch data
}

export const QuickTransferCard: React.FC<QuickTransferCardProps> = ({ userId }) => {
  const [user, setUser] = useState<UserProfile | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await axios.get(`https://bank-dashboard-6acc.onrender.com/users/${userId}`);
        setUser(response.data);
      } catch (err) {
        console.error("Failed to fetch user data:", err);
        setError("Failed to fetch user data. Please check the console for more details.");
      } finally {
        setLoading(false);
      }
    };

    fetchUserData();
  }, [userId]);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  if (!user) return <p>User not found.</p>;

  return (
    <div className="flex items-center space-x-4 p-2 bg-white rounded-lg shadow">
      <Image
        src={user.profilePicture}
        alt={`${user.username}'s profile picture`}
        width={44}
        height={44}
        className="rounded-full object-cover"
      />
      <p className="font-semibold text-sm md:text-base">{user.username}</p>
    </div>
  );
};

export default QuickTransferCard;
