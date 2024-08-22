import Card from "./(components)/card";
import { FaShieldHeart } from "react-icons/fa6";
import { GiShoppingBag } from "react-icons/gi";
import { AiFillSafetyCertificate } from "react-icons/ai";
import { FaHandHoldingDollar } from "react-icons/fa6";
import ServiceList from "./(components)/ServiceList";

const page = () => {
  return (
    <>
    <div className="bg-gray-100 px-[50px]">
    <div className="flex justify-between mt-10">
     <Card icon={FaShieldHeart} title="Life Insurance" sub_title="Unlimited Protection"/>
      <Card icon={GiShoppingBag} title="Shopping" sub_title="Buy.Think.Grow."/>
      <Card icon={AiFillSafetyCertificate} title="Safety" sub_title="We are your allies"/>
    </div>
    <p className="font-inter font-semibold text-[#343C6A] text-[22px] leading-[26.63px] my-5">Bank Services List</p>
     <ServiceList icon={FaHandHoldingDollar}/>
    </div>
   
      
    </>
  );
};

export default page;
