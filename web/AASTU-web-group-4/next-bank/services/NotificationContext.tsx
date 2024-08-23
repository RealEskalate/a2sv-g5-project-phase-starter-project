import React, { createContext, useContext, useState, useEffect, ReactNode, FC } from 'react';
import { format } from 'date-fns'; // Import date-fns for date formatting
import { Notification } from '@/types/index';
import { getAllTransactions } from '@/services/transactionfetch'; // Adjust the import path accordingly

import Cookies from "js-cookie";

const API_BASE_URL = "https://web-team-g4.onrender.com";
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

  // Fetch notifications from API and filter out already read notifications
  const fetchNotifications = async () => {
    try {
      const response = await getAllTransactions(0, 1000);
      console.log("Fetched response data:", response); // Check the full response structure

      if (response && response.data && Array.isArray(response.data.content)) {
        const readNotificationIds = JSON.parse(localStorage.getItem('readNotifications') || '[]');
        
        const formattedNotifications = response.data.content
          .filter((transaction: { transactionId: any }) => !readNotificationIds.includes(transaction.transactionId)) // Exclude read notifications
          .map((transaction: { transactionId: any; type: any; receiverUserName: any; senderUserName: any; date: any; amount: any; }) => ({
            id: transaction.transactionId,
            message: `${transaction.senderUserName} transferred you $${transaction.amount}`, // Show sender's name
            transactionId: transaction.transactionId,
            userId: transaction.senderUserName, // Use sender's name
            timestamp: transaction.date,
            formattedDate: format(new Date(transaction.date), 'MMM dd, yyyy'), // Format the date
            status: 'unread',
            isRead: false,
          }));
  
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
        status: 'read',
      }))
    );
  };

  useEffect(() => {
    fetchNotifications();
  }, []);

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
