import React from 'react';
import { IoMdNotificationsOutline } from 'react-icons/io';
import { useNotifications } from '@/services/NotificationContext';
import { DropdownMenu, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator } from '@/components/ui/dropdown-menu'; // Adjust import as needed

const NotificationBell: React.FC = () => {
  const { notifications, markAllAsRead, unreadCount } = useNotifications();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="relative cursor-pointer">
        <IoMdNotificationsOutline size={36} />
        {unreadCount > 0 && (
          <span className="absolute top-0 right-0 inline-flex items-center justify-center h-5 w-5 bg-red-500 text-white text-xs font-bold rounded-full">
            {unreadCount}
          </span>
        )}
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-80 max-h-80 overflow-y-auto scrollbar-thin scrollbar-thumb-gray-500 scrollbar-track-gray-200">
        <DropdownMenuLabel className="text-lg font-bold mb-2">Notifications</DropdownMenuLabel>
        <DropdownMenuSeparator />
        {notifications.length === 0 ? (
          <DropdownMenuItem>No recent notifications</DropdownMenuItem>
        ) : (
          <>
            {notifications.map(notification => (
              <DropdownMenuItem key={notification.id} className="flex justify-between items-center">
                <div>
                  <p className="text-sm">{notification.message}</p>
                  <span className="text-xs text-gray-500">{notification.formattedDate}</span>
                </div>
              </DropdownMenuItem>
            ))}
          </>
        )}
        <DropdownMenuSeparator />
        <DropdownMenuItem className="text-blue-500 hover:underline" onClick={() => markAllAsRead()}>
          Mark all as read
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default NotificationBell;
