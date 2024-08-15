import React from 'react';
import InfoboxCard from './InfoboxCard';
import { infoboxListItems } from './infoboxListItems';

const Infobox = () => {
  return (
    <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4 p-4">
      {infoboxListItems.map((item, index) => (
        <InfoboxCard key={index} name={item.name} icon={item.icon} value={item.value} />
      ))}
    </div>
  );
};

export default Infobox;
