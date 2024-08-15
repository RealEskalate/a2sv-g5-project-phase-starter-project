import React from "react";
import ServicesCard from "@/app/components/Card/ServicesCard";
import DescriptionCard from "@/app/components/Card/DescriptionCard";

const Services = () => {
  return (
    <div>
      <div className="flex gap-10 pt-10 pl-10">
        <ServicesCard
          img="/assets/lifeInsurance.svg"
          title="Life Insurance"
          desc="Unlimited Protection"
        />
        <ServicesCard
          img="/assets/shoppingBag.svg"
          title="Shopping"
          desc="Buy. Think. Grow"
        />
        <ServicesCard
          img="/assets/safety.svg"
          title="Safety"
          desc="We are your allies"
        />
      </div>
      <p className="font-semibold text-[22px] text-[#343C6A] p-10">
        Bank Services List
      </p>
      <div>
        <DescriptionCard
          img="/assets/loan.svg"
          title="Business loans"
          desc="It is a long established"
          colOne="Lorem Ipsum"
          descOne="Many publishing"
          colTwo="Lorem Ipsum"
          descTwo="Many publishing"
          colThree="Lorem Ipsum"
          descThree="Many publishing"
          btn="View Details"
          color="bg-pink-100"
        />
        <DescriptionCard
          img="/assets/accounts.svg"
          title="Checking accounts"
          desc="It is a long established"
          colOne="Lorem Ipsum"
          descOne="Many publishing"
          colTwo="Lorem Ipsum"
          descTwo="Many publishing"
          colThree="Lorem Ipsum"
          descThree="Many publishing"
          btn="View Details"
          color="bg-orange-100"
        />
        <DescriptionCard
          img="/assets/levels.svg"
          title="Saving accounts"
          desc="It is a long established"
          colOne="Lorem Ipsum"
          descOne="Many publishing"
          colTwo="Lorem Ipsum"
          descTwo="Many publishing"
          colThree="Lorem Ipsum"
          descThree="Many publishing"
          btn="View Details"
          color="bg-pink-100"
        />
        <DescriptionCard
          img="/assets/user.svg"
          title="Debit and credit cards"
          desc="It is a long established"
          colOne="Lorem Ipsum"
          descOne="Many publishing"
          colTwo="Lorem Ipsum"
          descTwo="Many publishing"
          colThree="Lorem Ipsum"
          descThree="Many publishing"
          btn="View Details"
          color="bg-blue-100"
        />
        <DescriptionCard
          img="/assets/safe.svg"
          title="Life insurance"
          desc="It is a long established"
          colOne="Lorem Ipsum"
          descOne="Many publishing"
          colTwo="Lorem Ipsum"
          descTwo="Many publishing"
          colThree="Lorem Ipsum"
          descThree="Many publishing"
          btn="View Details"
          color="bg-green-100"
        />
        <DescriptionCard
          img="/assets/loan.svg"
          title="Business loans"
          desc="It is a long established"
          colOne="Lorem Ipsum"
          descOne="Many publishing"
          colTwo="Lorem Ipsum"
          descTwo="Many publishing"
          colThree="Lorem Ipsum"
          descThree="Many publishing"
          btn="View Details"
          color="bg-pink-100"
        />
      </div>
    </div>
  );
};

export default Services;
