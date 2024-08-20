import React from "react";
import Link from "next/link";
import clsx from "clsx";

interface ServiceCardProps {
  icon: React.FC<React.SVGProps<SVGSVGElement>>;
  title: string;
  description: string;
  showDetailsLink?: boolean;
}

const ServiceCard: React.FC<ServiceCardProps> = ({
  icon: Icon,
  title,
  description,
  showDetailsLink = false,
}) => {
  return (
    <div
      className={clsx(
        "flex items-center p-4 shadow-lg rounded-lg bg-white",
        "lg:min-w-[350px] lg:h-[120px]", // Large screen width and height
        "sm:min-w-[230px] sm:h-[85px]" // Smaller screens width and height
      )}
    >
      <div className="flex items-center space-x-4">
        <Icon className="w-14 h-14" aria-hidden="true" />
        <div>
          <h3 className="text-lg font-semibold">{title}</h3>
          <p className="text-sm text-gray-500">{description}</p>
        </div>
      </div>
      {showDetailsLink && (
        <Link
          href="#"
          className="text-blue-500 text-sm font-semibold border border-blue-500 rounded-full px-2 py-1 ml-auto"
        >
          View Details
        </Link>
      )}
    </div>
  );
};

export default ServiceCard;
