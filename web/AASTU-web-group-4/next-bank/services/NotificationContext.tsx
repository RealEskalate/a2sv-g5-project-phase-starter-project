import React, { createContext, useContext, useState, useEffect, ReactNode, FC } from 'react';
import { format } from 'date-fns';
import { Notification, UserData } from '@/types/index';
import { getAllTransactions } from '@/services/transactionfetch'; 
import Cookies from "js-cookie";
import { getSession } from 'next-auth/react';
import { currentuser } from './userupdate';

const token = Cookies.get('accessToken');

type NotificationContextType = {
  notifications: Notification[];
  fetchNotifications: () => void;
  unreadCount: number;
  markAllAsRead: () => void;
};

const NotificationContext = createContext<NotificationContextType | undefined>(undefined);

export const NotificationProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const [notifications, setNotifications] = useState<Notification[]>([]);

  const [info, setInfo] = useState<UserData | null>(null)

  const fetchUserInfo = async() => {
    try {
      const data = await currentuser()
      setInfo(data.data)
      // console.log("data: ", data.data)
    } catch (error) {
      console.error("Error fetching user info:", error);
    }
  }
  
  const fetchNotifications = async () => {
    console.log("info: ", info)
    if (!info) {
      // Wait until user info is available
      console.log("User info not available yet.");
      return;
    }

    try {
      const response = await getAllTransactions(0, 1000);
      // console.log("Fetched response data:", response);
      
      
      

      if (response && response.data && Array.isArray(response.data.content)) {
        const readNotificationIds = JSON.parse(localStorage.getItem('readNotifications') || '[]');
        
        console.log(info)
        const currentUser = info.username
        // console.log("username: ", currentUser)
         // Assuming you store the current user's username in a cookie or other means
  
        // Process and format the notifications, adding a sequence number based on the index
        const formattedNotifications = response.data.content.map((transaction: { transactionId: any; type: any; receiverUserName: any; senderUserName: any; date: string; amount: any; }, index: number) => {
          const isSender = transaction.senderUserName === currentUser;
          const message = isSender
            ? `You have transferred $${transaction.amount} to ${transaction.receiverUserName}`
            : `${transaction.senderUserName} transferred you $${transaction.amount}`;
  
          return {
            id: transaction.transactionId,
            message,
            transactionId: transaction.transactionId,
            userId: transaction.senderUserName,
            timestamp: new Date(transaction.date).getTime(),
            formattedDate: format(new Date(transaction.date), 'MMM dd, yyyy'),
            isRead: readNotificationIds.includes(transaction.transactionId),
            sequence: index, // Use index as a sequence number
          };
        });
  
        // Sort notifications by timestamp (and sequence if needed)
        formattedNotifications.sort((a: { timestamp: number; sequence: number; }, b: { timestamp: number; sequence: number; }) => {
          // First, sort by timestamp in descending order
          if (b.timestamp !== a.timestamp) {
            return b.timestamp - a.timestamp;
          }
          // If timestamps are identical, sort by sequence number (descending)
          return b.sequence - a.sequence;
        });
  
        setNotifications(formattedNotifications);
      } else {
        console.error("Unexpected response structure:", response);
      }
    } catch (error) {
      console.error("Error fetching notifications:", error);
    }
  };
  
  
  
  
  

  const markAllAsRead = () => {
    const readNotificationIds = notifications.map(notification => notification.id);
    localStorage.setItem('readNotifications', JSON.stringify([...readNotificationIds, ...(JSON.parse(localStorage.getItem('readNotifications') || '[]'))]));
  
    setNotifications(prevNotifications =>
      prevNotifications.map(notification => ({
        ...notification,
        isRead: true,
      }))
    );
  };

  useEffect(() => {
    fetchUserInfo();
  }, []);
  useEffect(() => {
    fetchNotifications();
  }, [info]);

  const unreadCount = notifications.filter(notification => !notification.isRead).length;

  return (
    <NotificationContext.Provider value={{ notifications, fetchNotifications, unreadCount, markAllAsRead }}>
      {children}
    </NotificationContext.Provider>
  );
};

export const useNotifications = () => {
  const context = useContext(NotificationContext);
  if (!context) {
    throw new Error('useNotifications must be used within a NotificationProvider');
  }
  return context;
};

