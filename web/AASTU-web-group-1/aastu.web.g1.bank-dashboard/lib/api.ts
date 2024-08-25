import { getSession } from "next-auth/react";
import { TransactionProps,RandomBalanceHistory,UserResponse, CreditCardProps,QuickTransferProps, TransactionResponse } from "@/types";



export  async function getCreditCards(page: number, size: number) {
    try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    const res = await fetch(`${process.env.NEXT_PUBLIC_BASE_URL}/cards?page=${page}&size=${size}`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
      body: null,
    });
    if (!res.ok) {
      throw new Error("failed to get data");
    }
    const cards: CreditCardProps = await res.json();
    return cards;
  } catch (error) {
    console.error("An error occurred on card:", error);
  }
}

export async function getQuickTransfer(num: number) {
  try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/transactions/quick-transfers?number=${num}`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    );
    if (!res.ok) {
      throw new Error("failed to get data");
    }
    const account: QuickTransferProps = await res.json();
    return account.data;
  } catch (error) {
    console.error("An error occurred on card:", error);
  }
}

interface Props {
  type: string;
  description: string;
  amount: number;
  receiverUserName: string;
}

export async function addTransactions({
  type,
  description,
  amount,
  receiverUserName,
}: Props) {
  try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    console.log("add transaction",accessToken);
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/transactions`,
      {
        method: "POST",
        headers: {
          Authorization: `Bearer ${accessToken}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          type: type,
          description: description,
          amount: amount,
          receiverUserName: receiverUserName,
        }),
      }
    );
    if (!res.ok) {
      throw new Error("Failed to get data");
    }
  
    return true;
  } catch (error) {
    console.error("An error occurred on card:", error);
  }
}

export async function getCurrentUser() {
  try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/user/current`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    );
    if (!res.ok) {
      throw new Error("failed to get data");
    }

    const currentUser: UserResponse = await res.json();
    return currentUser.data;
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

export async function getExpenses(page: number, size: number) {
  try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/transactions/expenses?page=${page}&size=${size}`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    );
    if (!res.ok) {
      throw new Error("faild to fetch data");
    }
    const expenses: TransactionProps = await res.json();

    return expenses.data;
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

export async function getIncomes(page: number, size: number) {
  try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/transactions/incomes?page=${page}&size=${size}`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    );
    if (!res.ok) {
      throw new Error("faild to fetch data");
    }
    const incomes: TransactionProps = await res.json();

    return incomes.data;
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

export default async function getRandomBalance() {
  try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    console.log("random balance",accessToken);
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/transactions/random-balance-history?monthsBeforeFirstTransaction=7`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    );
    if (!res.ok) {
      throw new Error("faild to fetch data");
    }
    const balanceHistory: RandomBalanceHistory = await res.json();
    return balanceHistory.data;
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

export async function getallTransactions(page: number, size: number) {
  try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    console.log("get all transactions",accessToken);
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/transactions?page=${page}&size=${size}`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    );
    if (!res.ok) {
      throw new Error("faild to fetch data");
    }
    const transactions: TransactionProps = await res.json();

    return transactions.data;
  } catch (error) {
    console.log("An error occurred on transaction:", error);
  }
}

interface Prop {
  balance: number;
  cardHolder: string;
  expiryDate: string;
  passcode: string;
  cardType: string;
}

export async function postCards({
  cardHolder,
  expiryDate,
  passcode,
  cardType,
}: Prop) {
  try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    const res = await fetch(`${process.env.NEXT_PUBLIC_BASE_URL}/cards`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        balance: 300,
        cardHolder: cardHolder,
        expiryDate: expiryDate,
        passcode: passcode,
        cardType: cardType,
      }),
    });
    if (!res.ok) {
      throw new Error("Failed to get data");
    }
    const data = await res.json();
    alert("post Successful");
    console.log("res", data);
  } catch (error) {
    console.error("An error occurred on card:", error);
  }
}

export async function getbalance(){
  try{
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/transactions/balance-history`,
      {
        method: "GET",
        headers:{
          Authorization: `Bearer ${accessToken}`,
        }, 
        body: null,
      }
    );
    if (!res.ok){
      throw new Error("failed to get data");
    }

    const currentUser: UserResponse = await res.json();
    return currentUser.data;
  } catch(error) {
    console.log(error);
    console.log("An error occured:",error);
  }
}