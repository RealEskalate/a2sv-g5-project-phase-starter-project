import Image from 'next/image';
import React from 'react';
import { FaPiggyBank, FaHandshake, FaChartLine } from 'react-icons/fa'; // Import icons

const services = [
  { 
    title: 'Savings Accounts', 
    description: 'Flexible savings options to meet your goals.', 
    icon: FaPiggyBank 
  },
  { 
    title: 'Loans', 
    description: 'Affordable loan options for personal and business needs.', 
    icon: FaHandshake 
  },
  { 
    title: 'Investment Advice', 
    description: 'Expert advice to help you grow your wealth.', 
    icon: FaChartLine 
  }
];

const ServicesSection: React.FC = () => (
  <section className="py-16 bg-white flex">
    {/* Illustration Section */}
    <div className="hidden lg:w-1/2 lg:flex items-center justify-center">
      <Image 
        src="/Images/people-illustration.png" 
        alt="People illustration related to banking"
        width={700} 
        height={700} 
        className="max-w-full h-auto" 
      />
    </div>
    {/* Services Section */}
    <div className="w-[100%] lg:w-1/2 flex flex-col items-cenetr justify-center">
      <div className="container w-[100%] mx-auto px-6">
        <h2 className="text-3xl font-bold mb-8 text-left">Our Services</h2>
        <div className="space-y-8 w-[100%]  flex flex-col items-center" >
          {services.map((service, index) => (
            <div
              key={index}
              className="flex items-center bg-white p-6 rounded-lg shadow-lg transform transition-transform duration-300 hover:scale-105 hover:bg-blue-50"
            >
              {/* Icon */}
              <div className="p-4 bg-blue-100 rounded-full mr-6">
                <service.icon className="text-blue-600 h-8 w-8" />
              </div>
              {/* Text */}
              <div>
                <h3 className="text-2xl font-semibold mb-2 text-left">{service.title}</h3>
                <p className="text-gray-600 text-left">{service.description}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  </section>
);

export default ServicesSection;
