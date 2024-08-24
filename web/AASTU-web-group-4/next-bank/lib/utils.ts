/* eslint-disable no-prototype-builtins */
import { type ClassValue, clsx } from "clsx";
// import qs from "query-string";
import { twMerge } from "tailwind-merge";
// import { z } from "zod";
import { v4 as uuidv4 } from 'uuid';

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}



// src/data/generateDummyTransactions.ts


export interface Transaction {
  id: string;
  description: string;
  type: 'income' | 'expense';
  category: string;
  card: string;
  date: string;
  amount: number;
}

// Function to generate random date
const generateRandomDate = (): string => {
  const start = new Date(2023, 0, 1); // Jan 1, 2023
  const end = new Date();
  const date = new Date(start.getTime() + Math.random() * (end.getTime() - start.getTime()));
  return date.toLocaleString('en-US', { month: 'short', day: 'numeric', hour: 'numeric', minute: 'numeric', hour12: true });
};

// Dummy data generator
export const generateDummyTransactions = (count: number): Transaction[] => {
  const transactions: Transaction[] = [];

  const descriptions = [
    'Spotify Subscription', 
    'Mobile Services', 
    'Transfer from John Doe', 
    'Amazon Shopping', 
    'Electricity Bill'
  ];

  const categories = ['shopping', 'transfer', 'service'];
  const cards = ['Visa ****1234', 'Mastercard ****5678', 'Amex ****9101'];

  for (let i = 0; i < count; i++) {
    const isIncome = Math.random() > 0.5;
    transactions.push({
      id: uuidv4(),
      description: descriptions[Math.floor(Math.random() * descriptions.length)],
      type: isIncome ? 'income' : 'expense',
      category: categories[Math.floor(Math.random() * categories.length)],
      card: cards[Math.floor(Math.random() * cards.length)],
      date: generateRandomDate(),
      amount: isIncome ? +(Math.random() * 1000).toFixed(2) : -(Math.random() * 1000).toFixed(2),
    });
  }

  return transactions;
};

// Generate and export 100 dummy transactions
export const transactions = generateDummyTransactions(100);


export function convertDateToISOString(date: Date):string {
  if (date instanceof Date && !isNaN(date.getTime())) {
    return date.toISOString(); // Converts the date to ISO 8601 format
  } else {
    throw new Error("Invalid date");
  }
}

// FORMAT DATE TIME
// export const formatDateTime = (dateString: Date) => {
//   const dateTimeOptions: Intl.DateTimeFormatOptions = {
//     weekday: "short", // abbreviated weekday name (e.g., 'Mon')
//     month: "short", // abbreviated month name (e.g., 'Oct')
//     day: "numeric", // numeric day of the month (e.g., '25')
//     hour: "numeric", // numeric hour (e.g., '8')
//     minute: "numeric", // numeric minute (e.g., '30')
//     hour12: true, // use 12-hour clock (true) or 24-hour clock (false)
//   };

//   const dateDayOptions: Intl.DateTimeFormatOptions = {
//     weekday: "short", // abbreviated weekday name (e.g., 'Mon')
//     year: "numeric", // numeric year (e.g., '2023')
//     month: "2-digit", // abbreviated month name (e.g., 'Oct')
//     day: "2-digit", // numeric day of the month (e.g., '25')
//   };

//   const dateOptions: Intl.DateTimeFormatOptions = {
//     month: "short", // abbreviated month name (e.g., 'Oct')
//     year: "numeric", // numeric year (e.g., '2023')
//     day: "numeric", // numeric day of the month (e.g., '25')
//   };

//   const timeOptions: Intl.DateTimeFormatOptions = {
//     hour: "numeric", // numeric hour (e.g., '8')
//     minute: "numeric", // numeric minute (e.g., '30')
//     hour12: true, // use 12-hour clock (true) or 24-hour clock (false)
//   };

//   const formattedDateTime: string = new Date(dateString).toLocaleString(
//     "en-US",
//     dateTimeOptions
//   );

//   const formattedDateDay: string = new Date(dateString).toLocaleString(
//     "en-US",
//     dateDayOptions
//   );

//   const formattedDate: string = new Date(dateString).toLocaleString(
//     "en-US",
//     dateOptions
//   );

//   const formattedTime: string = new Date(dateString).toLocaleString(
//     "en-US",
//     timeOptions
//   );

//   return {
//     dateTime: formattedDateTime,
//     dateDay: formattedDateDay,
//     dateOnly: formattedDate,
//     timeOnly: formattedTime,
//   };
// };

// export function formatAmount(amount: number): string {
//   const formatter = new Intl.NumberFormat("en-US", {
//     style: "currency",
//     currency: "USD",
//     minimumFractionDigits: 2,
//   });

//   return formatter.format(amount);
// }

// export const parseStringify = (value: any) => JSON.parse(JSON.stringify(value));

// export const removeSpecialCharacters = (value: string) => {
//   return value.replace(/[^\w\s]/gi, "");
// };

// interface UrlQueryParams {
//   params: string;
//   key: string;
//   value: string;
// }

// export function formUrlQuery({ params, key, value }: UrlQueryParams) {
//   const currentUrl = qs.parse(params);

//   currentUrl[key] = value;

//   return qs.stringifyUrl(
//     {
//       url: window.location.pathname,
//       query: currentUrl,
//     },
//     { skipNull: true }
//   );
// }

// export function getAccountTypeColors(type: AccountTypes) {
//   switch (type) {
//     case "depository":
//       return {
//         bg: "bg-blue-25",
//         lightBg: "bg-blue-100",
//         title: "text-blue-900",
//         subText: "text-blue-700",
//       };

//     case "credit":
//       return {
//         bg: "bg-success-25",
//         lightBg: "bg-success-100",
//         title: "text-success-900",
//         subText: "text-success-700",
//       };

//     default:
//       return {
//         bg: "bg-green-25",
//         lightBg: "bg-green-100",
//         title: "text-green-900",
//         subText: "text-green-700",
//       };
//   }
// }

// export function countTransactionCategories(
//   transactions: Transaction[]
// ): CategoryCount[] {
//   const categoryCounts: { [category: string]: number } = {};
//   let totalCount = 0;

//   // Iterate over each transaction
//   transactions &&
//     transactions.forEach((transaction) => {
//       // Extract the category from the transaction
//       const category = transaction.category;

//       // If the category exists in the categoryCounts object, increment its count
//       if (categoryCounts.hasOwnProperty(category)) {
//         categoryCounts[category]++;
//       } else {
//         // Otherwise, initialize the count to 1
//         categoryCounts[category] = 1;
//       }

//       // Increment total count
//       totalCount++;
//     });

//   // Convert the categoryCounts object to an array of objects
//   const aggregatedCategories: CategoryCount[] = Object.keys(categoryCounts).map(
//     (category) => ({
//       name: category,
//       count: categoryCounts[category],
//       totalCount,
//     })
//   );

//   // Sort the aggregatedCategories array by count in descending order
//   aggregatedCategories.sort((a, b) => b.count - a.count);

//   return aggregatedCategories;
// }

// export function extractCustomerIdFromUrl(url: string) {
//   // Split the URL string by '/'
//   const parts = url.split("/");

//   // Extract the last part, which represents the customer ID
//   const customerId = parts[parts.length - 1];

//   return customerId;
// }

// export function encryptId(id: string) {
//   return btoa(id);
// }

// export function decryptId(id: string) {
//   return atob(id);
// }

// export const getTransactionStatus = (date: Date) => {
//   const today = new Date();
//   const twoDaysAgo = new Date(today);
//   twoDaysAgo.setDate(today.getDate() - 2);

//   return date > twoDaysAgo ? "Processing" : "Success";
// };

// export const authFormSchema = (type: string) => z.object({
//   // sign up
//   firstName: type === 'sign-in' ? z.string().optional() : z.string().min(3),
//   lastName: type === 'sign-in' ? z.string().optional() : z.string().min(3),
//   address1: type === 'sign-in' ? z.string().optional() : z.string().max(50),
//   city: type === 'sign-in' ? z.string().optional() : z.string().max(50),
//   state: type === 'sign-in' ? z.string().optional() : z.string().min(2).max(2),
//   postalCode: type === 'sign-in' ? z.string().optional() : z.string().min(3).max(6),
//   dateOfBirth: type === 'sign-in' ? z.string().optional() : z.string().min(3),
//   ssn: type === 'sign-in' ? z.string().optional() : z.string().min(3),
//   // both
//   email: z.string().email(),
//   password: z.string().min(8),
// })