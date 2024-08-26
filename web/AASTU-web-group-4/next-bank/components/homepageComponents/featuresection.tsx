import React from 'react';
import { FaMobileAlt, FaHeadset, FaMoneyBillAlt } from 'react-icons/fa'; // Importing icons

const features = [
  { title: 'Online Banking', description: 'Manage your money on the go', icon: FaMobileAlt },
  { title: '24/7 Support', description: 'We are here for you anytime', icon: FaHeadset },
  { title: 'Low Fees', description: 'Experience transparent and low fees', icon: FaMoneyBillAlt }
];

const FeaturesSection: React.FC = () => (
  <section className="py-16 bg-gray-100">
    <div className="container mx-auto text-center">
      <h2 className="text-3xl font-bold mb-8">Our Features</h2>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
        {features.map((feature, index) => (
          <div
            key={index}
            className="bg-white p-6 rounded-lg shadow-lg transform transition-transform duration-300 hover:scale-105 animate__animated animate__fadeIn"
          >
            <feature.icon className="mx-auto mb-4 h-24 w-24 text-blue-400" /> {/* Dynamic icon */}
            <h3 className="text-xl font-semibold mb-2">{feature.title}</h3>
            <p className="text-gray-700">{feature.description}</p>
          </div>
        ))}
      </div>
    </div>
  </section>
);

export default FeaturesSection;
