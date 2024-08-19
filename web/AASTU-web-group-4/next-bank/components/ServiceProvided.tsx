import React from "react";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
} from "@/components/ui/carousel";
import LifeInsuranceIcon from "@/public/icons/LifeInsuranceIcon";
import ShoppingIcon from "@/public/icons/ShoppingIcon";
import SafetyIcon from "@/public/icons/SafetyIcon";

const services = [
  {
    title: "Life Insurance",
    description: "Unlimited protection",
    icon: LifeInsuranceIcon,
  },
  {
    title: "Shopping",
    description: "Buy, Think, Grow",
    icon: ShoppingIcon,
  },
  {
    title: "Safety",
    description: "We are your allies",
    icon: SafetyIcon,
  },
];

const ServiceProvided: React.FC = () => {
  return (
    <div className="lg:max-w-[1110px] lg:mx-auto">
      <Carousel className="lg:hidden">
        <CarouselContent className="py-4">
          {services.map((service, index) => (
            <CarouselItem
              key={index}
              className="w-[230px] h-[85px] mx-auto mr-4 flex-none"
            >
              <div className="shadow-lg p-4 rounded-md flex items-center h-full">
                <service.icon className="w-13 h-13 mr-4" aria-hidden="true" />
                <div>
                  <h3 className="text-[14px] font-semibold">{service.title}</h3>
                  <p className="text-[12px] text-gray-500">
                    {service.description}
                  </p>
                </div>
              </div>
            </CarouselItem>
          ))}
        </CarouselContent>
      </Carousel>

      {/* For Large Screens */}
      <div className="hidden lg:flex gap-8">
        {services.map((service, index) => (
          <div
            key={index}
            className="w-[350px] h-[120px] shadow-lg p-4 rounded-md flex items-center"
          >
            <service.icon className="w-130 h-130 mr-4" aria-hidden="true" />
            <div>
              <h3 className="text-[20px] font-semibold">{service.title}</h3>
              <p className="text-[16px] text-gray-500">{service.description}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ServiceProvided;
