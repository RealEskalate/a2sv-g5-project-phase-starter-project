import React from 'react';
import ServiceCard from './components/ServiceCard';
import ServiceListCard from './components/ServiceListCard';

const HomePage = () => {
  const categories = [
    { iconSrc: '/insuranceicon.svg', title: 'Life Insurance', subtitle: 'Unlimited protection', bgcolor: 'bg-blue-100' },
    { iconSrc: '/shopicon.svg', title: 'Shopping', subtitle: 'Buy. Think. Grow.', bgcolor: 'bg-yellow-100' },
    { iconSrc: '/safetyicon.svg', title: 'Safety', subtitle: 'We are your allies', bgcolor: 'bg-green-100' },
  ];

  const services = [
    {
      iconSrc: '/businessloan.svg',
      mainTitle: 'Business loans',
      mainSubtitle: 'It is a long established',
      details: [
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
      ],
      buttonText: 'View Details',
      bgColor: 'bg-pink-100',
    },
    {
      iconSrc: '/checkaccount.svg',
      mainTitle: 'Checking accounts',
      mainSubtitle: 'It is a long established',
      details: [
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
      ],
      buttonText: 'View Details',
      bgColor: 'bg-yellow-100',
    },
    {
      iconSrc: '/saveaccount.svg',
      mainTitle: 'Savings accounts',
      mainSubtitle: 'It is a long established',
      details: [
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
      ],
      buttonText: 'View Details',
      bgColor: 'bg-pink-100',
    },
    {
      iconSrc: '/debitcredit.svg',
      mainTitle: 'Debit and credit cards',
      mainSubtitle: 'It is a long established',
      details: [
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
      ],
      buttonText: 'View Details',
      bgColor: 'bg-blue-100',
    },
    {
      iconSrc: '/insuranceicon.svg',
      mainTitle: 'Life Insurance',
      mainSubtitle: 'It is a long established',
      details: [
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
      ],
      buttonText: 'View Details',
      bgColor: 'bg-green-100',
    },
    {
      iconSrc: '/businessloan.svg',
      mainTitle: 'Business Loans',
      mainSubtitle: 'It is a long established',
      details: [
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
        { title: 'Lorem Ipsum', subtitle: 'Many publishing' },
      ],
      buttonText: 'View Details',
      bgColor: 'bg-pink-100',
    },

  ];

  return (
    <div className="container mx-auto px-4 py-0">
      <div className="flex overflow-x-auto gap-10 mb-6 whitespace-nowrap pt-10 pb-10 mr-10 ml-10 mt-5 mb-5">
        {categories.map((category, index) => (
          
          <ServiceCard
            key={index}
            iconSrc={category.iconSrc}
            title={category.title}
            subtitle={category.subtitle}
            bgColor={category.bgcolor}
            
          />
        ))}
      </div>
      
      <h1 className="text-xl font-semibold text-gray-800 mb-6">Bank Services List</h1>
      
      <div className="space-y-4"> {/* Adding vertical space between ServiceListCards */}
        {services.map((service, index) => (
          <ServiceListCard
            key={index}
            iconSrc={service.iconSrc}
            mainTitle={service.mainTitle}
            mainSubtitle={service.mainSubtitle}
            details={service.details}
            buttonText={service.buttonText}
            bgColor={service.bgColor}
          />
        ))}
      </div>
    </div>
  );
};

export default HomePage;
