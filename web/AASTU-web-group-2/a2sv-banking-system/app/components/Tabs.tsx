import React from 'react';

interface TabsProps {
  tabs: string[]; 
  activeTab: string;
  onTabChange: (tab: string) => void;
}

const Tabs: React.FC<TabsProps> = ({ tabs, activeTab, onTabChange }) => {
  return (
    <div className="flex space-x-12 border-b pb-4">
      {tabs.map((tab) => (
        <button
          key={tab}
          className={`font-black ${activeTab === tab ? 'text-[#1814F3] border-b-2 border-[#1814F3]' : 'text-[#718EBF] dark:text-[#9faaeb]' }`}
          onClick={() => onTabChange(tab)}
        >
          {tab}
        </button>
      ))}
    </div>
  );
};

export default Tabs;
