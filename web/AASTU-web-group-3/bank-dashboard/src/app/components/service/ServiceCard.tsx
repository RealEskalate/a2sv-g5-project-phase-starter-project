"use client"
import Link from "next/link";
import Image from "next/image";
import { ServiceType } from "@/types/serviceCard";
import { menuItems } from "@/../../public/Icons";
import { useDispatch, useSelector } from "react-redux";
import { useGetAllServiceQuery } from "@/lib/redux/api/serviceApi";
import {
  setService,
  setLoading,
  setError,
} from "@/lib/redux/slices/serviceSlice";
import { RootState } from "@/lib/redux/store";
import { useEffect } from "react";
import Loading from "@/app/loading";

export default function ServiceCard() {


  const dispatch = useDispatch();
  const { service, loading, error } = useSelector(
    (state: RootState) => state.service
  );

  console.log(service)
  const { data, isLoading, isError } = useGetAllServiceQuery({
    size: 10,
    page: 0,
  });

  useEffect(() => {
    dispatch(setLoading(isLoading));

    if (data) {
      dispatch(setService(data.data.content));
    }

    if (isError) {
      dispatch(setError("Error loading transactions"));
    }
  }, [data, isLoading, isError, dispatch]);

  if (loading) return <Loading />;
  if (error) return <div>{error}</div>;

  return (
    <>
    
      {service.map((card, index: number) => (
        <div key={index} className="body flex  md:w-auto w-auto h-auto p-2 border-[1px] rounded-[10px] m-2 bg-white">
          <div className="flex items-center rounded-2xl px-5 bg-[#FFE0EB]">
            <Image width={18} height={18} src={card.icon} alt="aastu" />
          </div>
          <div className="right w-full flex justify-between items-center p-2">
            <div className="md:w-1/4 flex-shrink-0">
              <div className="font-normal ">{card.name}</div>
              <div className="font-normal h-5 flex-grow max-w-xs truncate text-xs text-[#718EBF]">
                {card.details}
              </div>
            </div>
            <div className="hidden md:block md:w-1/6">
              <div className="font-medium text-sm md:text-[12px]">Status</div>
              <div className="font-normal text-xs text-[#718EBF]">{card.status}</div>
            </div>
            <div className="hidden md:block md:w-1/6">
              <div className="font-medium text-sm md:text-[12px]"> Type</div>
              <div className="font-normal text-xs text-[#718EBF]">{card.type}</div>
            </div>
            <div className="hidden md:block md:w-1/6">
              <div className="font-medium text-sm md:text-[12px]">
                Number of users
              </div>
              <div className="font-normal text-xs text-[#718EBF]">
                {card.numberOfUsers}
              </div>
            </div>
            {/* <div className="md:px-4 md:py-1 md:border  md:border-[#718EBF] md:rounded-full hover:border-[#1814F3] text-center"> */}
            <Link
              href={"/services"}
              className="md:px-4 md:py-2 md:border  md:border-[#718EBF] md:rounded-full hover:border-[#1814F3] text-center font-normal text-[11px] text-[#1814F3] md:text-[#718EBF] hover:text-[#1814F3]"
            >
              View Details
            </Link>
          </div>
        </div>
      ))}
    </>
  );
}
