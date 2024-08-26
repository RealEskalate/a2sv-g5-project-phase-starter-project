import React from "react";
import { Item } from "../../lastTransaction/lastTransactionItems";
interface NotificationsProps {
  notifications: Item[];
}

const Notification = ({ notifications }: NotificationsProps) => {
  return (
    <div className="absolute right-0 mt-2 w-64 rounded-md shadow-lg py-2 z-50 bg-white">
      <p className="px-4 py-2 text-sm font-semibold text-gray-700">
        Notifications
      </p>
      <ul className="max-h-48 overflow-auto">
        {notifications.map((notification, index) => (
          <li
            key={index}
            className="px-4 py-2 text-sm text-gray-600 hover:bg-gray-100 cursor-pointer"
          >
            {index+1  + ". " + notification.amount + " "+ notification.type}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Notification;
