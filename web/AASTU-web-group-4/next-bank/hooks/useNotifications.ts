// // src/hooks/useNotifications.ts
// import { useState, useEffect } from 'react';
// import { getFirestore, collection, query, onSnapshot } from 'firebase/firestore';
// import { Notification } from '@/types'; // Adjust import path as needed

// const db = getFirestore(); // Ensure Firestore is initialized

// const useNotifications = () => {
//   const [notifications, setNotifications] = useState<Notification[]>([]);
//   const [error, setError] = useState<string | null>(null);

//   useEffect(() => {
//     const q = query(collection(db, 'notifications'));

//     const unsubscribe = onSnapshot(
//       q,
//       (querySnapshot) => {
//         const newNotifications: Notification[] = [];
//         querySnapshot.forEach((doc) => {
//           const data = doc.data();
//           const notification: Notification = {
//             id: doc.id,
//             ...data as Omit<Notification, 'id'>
//           };
//           newNotifications.push(notification);
//         });
//         setNotifications(newNotifications);
//         setError(null); // Clear any previous errors
//       },
//       (error) => {
//         console.error('Error fetching notifications:', error);
//         setError('Failed to fetch notifications.');
//       }
//     );

//     return () => unsubscribe();
//   }, []);

//   return { notifications, error };
// };

// export default useNotifications;
