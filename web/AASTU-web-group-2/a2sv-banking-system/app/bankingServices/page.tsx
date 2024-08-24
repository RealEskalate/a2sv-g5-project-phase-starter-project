'use client'
import React, { useEffect, useState } from 'react'
import { useRouter } from "next/navigation";
import InformationCard from './components/InformationCard'
import BankServiceList from './components/BankServiceList'
import { getSession } from 'next-auth/react';
import { IconType } from 'react-icons';
import Refresh from "../api/auth/[...nextauth]/token/RefreshToken";
import { getBankServices } from '@/lib/api/bankServiceControl';

const shimmerClass = 'bg-gray-200 animate-pulse';

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
  console.log(access_token)
  useEffect(() => {
    const addingData = async () => {
      if (!access_token) return;
      if (access_token) {
        const bankServices = await getBankServices(access_token, 0, 100);
        console.log("Fetching Completed", bankServices.data.content);
        setBankServices(bankServices.data.content); // Set the content array
      }
    };
    addingData();
  }, [access_token]);

  if (loading) {
    // Render shimmer effect while loading
    return (
      <div className="flex-col bg-gray-100 p-4">
        <div className="flex mx-5 my-4 gap-4 overflow-x-auto">
          <div className={`${shimmerClass} w-40 h-10 rounded-md mb-4`} />
          <div className={`${shimmerClass} w-full h-8 rounded-md mb-4`} />
          <div className={`${shimmerClass} w-full h-32 rounded-md`} />
        </div>
      </div>
    );
  }




  
  if (!session) {
    router.push(`./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
    return null;
  }

  const getBackgroundColor = (index: number) => {
    const colors = ['bg-[#FFE0EB]', 'bg-[#E0F7FA]', 'bg-[#FFF9C4]'];
    return colors[index % colors.length];
  };

  return (
    <div className='flex-col bg-[#f5f7fa]'>



      <div className='flex mx-5 my-4 rounded-3xl gap-4 overflow-x-auto [&::-webkit-scrollbar]:hidden'>
      <InformationCard
        logoBgColor="#e7edff"
        logo={
              <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
              <g clip-path="url(#clip0_163_357)">
              <path d="M18.2112 2.72841C15.6138 2.72841 13.6258 2.07743 12.4154 1.53132C11.0997 0.937663 10.3902 0.335245 10.3843 0.330201L10.001 0L9.61716 0.329224C9.61016 0.335245 8.90061 0.937702 7.58493 1.53132C6.37462 2.07743 4.38661 2.72841 1.78914 2.72841H1.20264V10.7793L10.0002 20L18.7977 10.7874V2.72841H18.2112ZM16.4517 8.20503C16.4517 9.11454 16.0983 9.9697 15.4564 10.6137L10.0002 16.2817L4.54395 10.6137C3.90204 9.96966 3.54865 9.1145 3.54865 8.20499V7.90008C3.54865 6.01862 5.07935 4.48788 6.96085 4.48788C7.80522 4.48788 8.61588 4.79897 9.24348 5.36381L9.97892 6.0257L10.5173 5.48729C11.1618 4.8428 12.0187 4.48785 12.9301 4.48785H13.0395C14.921 4.48785 16.4517 6.01854 16.4517 7.90004V8.20503H16.4517Z" fill="#396AFF"/>
              <path d="M13.0395 5.66093H12.9302C12.3321 5.66093 11.7698 5.89385 11.3468 6.31676L10.0215 7.6421L8.45884 6.2357C8.04696 5.86503 7.515 5.66089 6.96087 5.66089C5.72616 5.66089 4.72168 6.66537 4.72168 7.90008V8.20498C4.72168 8.8031 4.9546 9.36544 5.37751 9.78835L5.38533 9.79628L10.0002 14.5902L14.6229 9.78831C15.0458 9.3654 15.2787 8.80306 15.2787 8.20494V7.90004C15.2787 6.66545 14.2742 5.66093 13.0395 5.66093ZM11.7597 10.9395H10.5867V12.1125H9.41371V10.9395H8.2407V9.76645H9.41371V8.59344H10.5867V9.76645H11.7597V10.9395Z" fill="#396AFF"/>
              </g>
              <defs>
              <clipPath id="clip0_163_357">
              <rect width="20" height="20" fill="white"/>
              </clipPath>
              </defs>
              </svg>
        }
        title="Life Insurance"
        description="Unlimited protection"
        cardBgColor="bg-[#ffffff]"
      />
            <InformationCard
                  logoBgColor="#fff5d9"
                  logo={
          <svg width="16" height="20" viewBox="0 0 16 20" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M15.3125 17.275L14.2463 5.56875C14.2175 5.24625 13.9475 5 13.6238 5H11.7488V3.75C11.7488 2.745 11.3588 1.8025 10.6525 1.09625C9.95752 0.4 8.99002 0 7.99877 0C5.93127 0 4.24877 1.6825 4.24877 3.75V5H2.37377C2.05002 5 1.78002 5.24625 1.75127 5.56875L0.687524 17.2738C0.623774 17.9725 0.858775 18.6688 1.33127 19.1863C1.80377 19.7038 2.47627 20 3.17752 20H12.8213C13.5213 20 14.1938 19.7038 14.6663 19.1875C15.14 18.67 15.3738 17.9725 15.3125 17.275ZM10.4988 5H5.49877V3.75C5.49877 2.37125 6.62002 1.25 7.99877 1.25C8.66127 1.25 9.30627 1.515 9.76877 1.97875C10.24 2.45 10.4988 3.07875 10.4988 3.75V5Z" fill="#FFBB38"/>
          </svg>

        }
        title="Shopping"
        description="Buy. Think. Grow"
        cardBgColor="bg-[#ffffff]"
      />
            <InformationCard
        logoBgColor="#dcfaf8"
        logo={
          <svg width="18" height="20" viewBox="0 0 18 20" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M16.6553 13.3559C16.1178 14.8128 15.3048 16.0796 14.2387 17.1207C13.0251 18.3057 11.4361 19.2473 9.5156 19.919C9.45258 19.941 9.38681 19.959 9.32074 19.9722C9.23346 19.9895 9.14466 19.9988 9.05661 20H9.03937C8.94553 20 8.85123 19.9905 8.75769 19.9722C8.69162 19.959 8.62677 19.941 8.56406 19.9194C6.6413 19.2488 5.05026 18.3077 3.83551 17.1227C2.76892 16.0815 1.95608 14.8155 1.41928 13.3586C0.443171 10.7097 0.498713 7.79159 0.543421 5.44662L0.544184 5.41061C0.553187 5.21697 0.558985 5.01357 0.562189 4.78896C0.578516 3.68621 1.45529 2.77389 2.55819 2.71239C4.85769 2.58407 6.63657 1.8341 8.1565 0.35262L8.16977 0.340413C8.42215 0.108937 8.74014 -0.00458834 9.05661 0.00014191C9.36179 0.00410921 9.66574 0.117482 9.90912 0.340413L9.92209 0.35262C11.4423 1.8341 13.2212 2.58407 15.5207 2.71239C16.6236 2.77389 17.5004 3.68621 17.5167 4.78896C17.5199 5.0151 17.5257 5.21819 17.5347 5.41061L17.5352 5.42587C17.5797 7.77526 17.635 10.6992 16.6553 13.3559Z" fill="#16DBCC"/>
          <path d="M16.6554 13.356C16.1178 14.8129 15.3048 16.0797 14.2387 17.1208C13.0251 18.3058 11.4361 19.2474 9.51562 19.9191C9.45261 19.9411 9.38684 19.9591 9.32077 19.9723C9.23349 19.9896 9.14468 19.9989 9.05664 20.0001V0.000244141C9.36182 0.00421144 9.66577 0.117584 9.90915 0.340515L9.92212 0.352722C11.4424 1.8342 13.2212 2.58417 15.5207 2.7125C16.6236 2.77399 17.5004 3.68631 17.5167 4.78907C17.5199 5.0152 17.5257 5.2183 17.5347 5.41071L17.5352 5.42597C17.5797 7.77537 17.635 10.6993 16.6554 13.356Z" fill="#16DBCC"/>
          <path d="M14.0237 10.0001C14.0237 12.7425 11.797 14.9749 9.05682 14.9847H9.03927C6.29101 14.9847 4.05469 12.7485 4.05469 10.0001C4.05469 7.2518 6.29101 5.01562 9.03927 5.01562H9.05682C11.797 5.02539 14.0237 7.25775 14.0237 10.0001Z" fill="white"/>
          <path d="M12.1015 9.19753L9.07623 11.8814L8.42252 12.4613C8.26809 12.5983 8.06554 12.6667 7.86319 12.6667C7.66064 12.6667 7.4583 12.5983 7.30366 12.4613L5.89815 11.214C5.58929 10.94 5.58929 10.4962 5.89815 10.2221C6.2066 9.94812 6.70753 9.94812 7.01639 10.2221L7.86319 10.9732L10.9833 8.20561C11.2921 7.93146 11.7931 7.93146 12.1015 8.20561C12.4104 8.47959 12.4104 8.92392 12.1015 9.19753Z" fill="#16DBCC"/>
          </svg>
        }
        title="Safety"
        description="We are your allies"
        cardBgColor="bg-[#ffffff]"
      />
      </div>
      <h1 className='text-[#343C6A] font-semibold mx-5 my-4 text-2xl md:font-bold'>Bank Services List</h1>
      <div className='flex-col gap-5'>
        {bankInfo.map((item, index) => (
          <BankServiceList
            key={item.id}
            logoBgColor={getBackgroundColor(index)}
            logoSvg={(
              <svg className="w-8 h-8" viewBox="0 0 24 24">
                <image
                  href={item.icon}
                  x="3" // Adjust this value to center the image horizontally
                  y="3" // Adjust this value to center the image vertically
                  width="18" // Smaller width
                  height="18" // Smaller height
                />
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
