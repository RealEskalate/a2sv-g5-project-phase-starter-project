import Link from "next/link";
import React from "react";

const Card = ({
  title,
  button,
  link,
  children,
}: {
  title: string;
  button: string;
  link: string;
  children: React.ReactNode;
}) => {
  return (
    <div className="flex flex-col lg:w-[730px] lg:h-[300px] md:w-[487px] md:h-[299px] h-[254]">
      <div className="flex justify-between">
        <p className="font-semibold my-3 text-[18px] md:text-[18px] lg:text[22px] text-[#343C6A]">{title}</p>
        <Link href={link}>
          <p className="font-semibold my-3 text[12px] text-[#343C6A] text-right pr-3">
            {button}
          </p>
        </Link>
      </div>
      {children}
    </div>
  );
};

export default Card;
