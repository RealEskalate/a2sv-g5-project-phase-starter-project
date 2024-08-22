import Link from "next/link";
import React from "react";

const CardForCreditCards = ({
  title,
  button,
  link,
  children,
  className = "",
}: {
  title: string;
  button: string;
  link: string;
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <div className={`${className} mt-3 `}>
      <div className="flex justify-between items-center mb-4">
        <p className="font-semibold text-lg md:text-xl lg:text-2xl text-[#343C6A]">
          {title}
        </p>
        <Link
          href={link}
          className="font-semibold text-sm text-[#343C6A] hover:underline"
        >
          {button}
        </Link>
      </div>
      <div className="flex flex-wrap gap-4 overflow-x-auto scrollbar-hide">{children}</div>
    </div>
  );
};

export default CardForCreditCards;
