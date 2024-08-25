import { getSession } from "next-auth/react";
import { FormValues } from "./AddCardModal";

export function formatDateString(dateString: string) {
  const date = new Date(dateString);
  const month = (date.getMonth() + 1).toString().padStart(2, "0");
  const year = date.getFullYear().toString().slice(-2);
  return `${month}/${year}`;
}



export function Dateformat(dateString: string) {
  const date = new Date(dateString);
  const month = (date.getMonth() + 1).toString().padStart(2, "0");
  const year = date.getFullYear().toString().slice(-2);
  return `${year}-${month}-${date.getDate()}`;
}


export 
function formatMonth(dateString: string) {
  const [year, month] = dateString.split("-");
  const monthNames = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];
  return monthNames[parseInt(month) - 1];
}


export function maskCardNumber(cardNumber: string) {
  // Ensure the card number is a string
  const strCardNumber = cardNumber.toString();

  const length = strCardNumber.length;

  if (length <= 8) {
    const firstFour = strCardNumber.slice(0, 4);
    const lastFour = strCardNumber.slice(-4);
    const maskedSection = "****";
    return `${firstFour} ${maskedSection} ${lastFour}`;
  }

  const firstFour = strCardNumber.slice(0, 4);
  const lastFour = strCardNumber.slice(-4);
  const maskedSection = strCardNumber
    .slice(4, -4)
    .replace(/./g, "*")
    .replace(/(.{4})/g, "$1 ");

  return `${firstFour} ${maskedSection.trim()} ${lastFour}`;
}
 

export async function AddCard(data:FormValues){
   try {

                 const session = await getSession();
                 const accessToken = session?.user.accessToken;
                 const res = await fetch(
                   `${process.env.NEXT_PUBLIC_BASE_URL}/cards`,
                   {
                     method: "POST",
                     headers: {
                       Authorization: `Bearer ${accessToken}`,
                       "Content-Type": "application/json",
                     },
                     body: JSON.stringify({
                       balance: 300,
                       cardHolder: data.cardHolder,
                       expiryDate: data.expiryDate,
                       passcode: data.passcode,
                       cardType: data.cardType,
                     }),
                   }
                 );if (!res.ok) {
                   throw new Error("Failed to get data");} return true;
} catch (error) {
                 console.error("Failed to submit form:", error);
               }
                 
          

}