import Link from "next/link";
import React from "react";

const CardForCreditCards = ({
  title,
  button,
  link,
  children,
  className = "flex flex-col lg:w-[730px] lg:h-[300px] md:w-[487px] md:h-[299px] h-[254px] max-md:w-[550px]",
}: {
  title: string;
  button: string;
  link: string;
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <div className={`${className} max-md:overflow-x-auto`}> 
      <div className="flex justify-between">
        <p className="font-semibold my-3 text-[18px] md:text-[18px] lg:text-[22px] text-[#343C6A]">{title}</p>
        <Link href={link}>
          <p className="font-semibold my-3 text-[12px] text-[#343C6A] text-right pr-3">
            {button}
          </p>
        </Link>
      </div>
      <div className="flex justify-between max-md:overflow-x-auto">
        {children}
      </div>
    </div>
  );
};

export default CardForCreditCards;
