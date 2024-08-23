'use client'
import React, { useEffect, useState } from 'react'
import { useRouter } from "next/navigation";
import InformationCard from './components/InformationCard'
import BankServiceList from './components/BankServiceList'
import { getSession } from 'next-auth/react';
import { IconType } from 'react-icons';
import Refresh from "../api/auth/[...nextauth]/token/RefreshToken";
import { getBankServices } from '@/lib/api/bankServiceControl';

type infoType = {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}

type DataItem = {
  heading: string;
  text: string;
  headingStyle: string;
  dataStyle: string;
};

type Column = {
  icon: IconType;
  iconStyle: string;
  data: DataItem[];
};

type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};

type SessionDataType = {
  user: Data;
};

const Page = () => {
  const [session, setSession] = useState<Data | null>(null);
  const [access_token, setAccess_token] = useState("");
  const router = useRouter();
  const [loading, setLoading] = useState(true);
  const [bankInfo, setBankServices] = useState<infoType[]>([]);
  
  useEffect(() => {
    const fetchSession = async () => {
      try {
        const sessionData = (await getSession()) as SessionDataType | null;
        setAccess_token(await Refresh());
        if (sessionData && sessionData.user) {
          setSession(sessionData.user);
        } else {
          router.push(`./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
        }
      } catch (error) {
        console.error("Error fetching session:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchSession();
  }, [router]);

  useEffect(() => {
    const addingData = async () => {
      if (!access_token) return;
      if (access_token) {
        const bankServices = await getBankServices( 0, 100, access_token);
        console.log("Fetching Completed", bankServices.data);
        setBankServices(bankServices.data); // Set the content array
      }
    };
    addingData();
  }, [access_token]);

  if (loading) return null;

  if (!session) {
    router.push(`./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
    return null;
  }

  return (
    <div className='flex-col bg-[#f5f7fa]'>
      <div className='flex mx-5 my-4 rounded-3xl gap-4 overflow-x-auto [&::-webkit-scrollbar]:hidden'>
      </div>
      <h1 className='text-[#343C6A] font-semibold mx-5 my-4 text-xl md:font-bold'>Bank Services List</h1>
      <div className='flex-col gap-5'>
        {bankInfo.map((item, index) => (
          <BankServiceList
            key={index}
            logoBgColor="bg-[#FFE0EB]"
            logoSvg={(
          <svg className="w-8 h-8 text-blue-500 items-center" viewBox="0 0 24 24">
            <image href={item.icon} width="24" height="24" />
          </svg>

            )}
            serviceName={item.name}
            serviceDescription={item.details}
            additionalServices={[
              { name: `Users: ${item.numberOfUsers}`, description: "" },
              { name: `Type: ${item.type}`, description: "" },
              { name: `Status: ${item.status}`, description: ""},
            ]}
            viewDetailsLink={`https://example.com/details/${index}`}
          />
        ))}
      </div>
    </div>
  );
}

export default Page;
