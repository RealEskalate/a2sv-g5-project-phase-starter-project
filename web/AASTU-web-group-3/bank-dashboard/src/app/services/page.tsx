import { ServiceType,TopPageCardType } from "@/types/serviceCard";
import ServiceCard from "@/app/components/service/ServiceCard";
import TopPageCard from "@/components/ui/TopPageCard"
import {safety,shooping,lifeinsurance } from "@/../../public/Icons";


  
const Service = () =>{
    const TopCard:TopPageCardType[] =[
        {
            title: "Life Insurance",
            subTitle: "Unlimited protection",
            svgIcon:lifeinsurance,
            bgColor: "bg-[#E7EDFF]"
        },
        {
            title: "Shopping",
            subTitle: "Buy. Think. Grow",
            svgIcon:shooping,
            bgColor: "bg-orange-100"
        },
        {
            title: "Safety",
            subTitle: "We are your allies",
            svgIcon:safety,
            bgColor: "bg-[#DCFAF8]"
        },
    ]
      
    return (
      <div className="w-full ">
         <div className="w-11/12 m-3 flex gap-5 md:flex-row overflow-x-auto lg:overflow-visible  overflow-y-hidden no-scrollbar items-center md:justify-between">
         {TopCard.map((items,index) =>(
            <div key={index} className="w-3/4 md:w-1/2 lg:w-1/3 flex-shrink-0  ">
            <TopPageCard svgIcon={items.svgIcon} title = {items.title} subTitle={items.subTitle} bgColor={items.bgColor}/>
            </div>

         ))

         }
        </div>
        <div className="p-4 font-semibold text-[16px] text-xl text-[#343C6A] mx-2">Bank Services List</div>
          <ServiceCard/>
      </div>
    )
}

export default Service;