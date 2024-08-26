"use client"

import React, { useState } from 'react';
import { Accordion, AccordionSummary, AccordionDetails, Typography } from '@mui/material';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';

type FAQItem = {
  question: string;
  answer: string;
};

const faqItems: FAQItem[] = [
  { question: 'What services does the bank offer?', answer: 'We provide personal and business banking solutions including savings, loans, and investment advice.' },
  { question: 'How can I open a new account?', answer: 'You can open an account online or visit any of our branches with valid identification.' },
  { question: 'Is there a minimum balance requirement?', answer: 'Yes, the minimum balance varies by account type. Please refer to our account comparison page for details.' },
  { question: 'How do I reset my online banking password?', answer: 'You can reset your password via the “Forgot Password” link on the login page, or by contacting our support team.' },
  { question: 'What are your branch working hours?', answer: 'Our branches are open from 9:00 AM to 5:00 PM, Monday through Friday.' },
  { question: 'Are there any fees for international transfers?', answer: 'Yes, fees may apply for international transfers depending on the destination and the amount being transferred.' },
  { question: 'How secure is online banking?', answer: 'We use industry-standard encryption and security measures to ensure your information is safe when using online banking.' },
];

const FAQSection: React.FC = () => {
  const [expanded, setExpanded] = useState<number | false>(false);

  const handleChange = (panel: number) => (event: React.SyntheticEvent, isExpanded: boolean) => {
    setExpanded(isExpanded ? panel : false);
  };

  return (
    <div className='py-10' style={{ maxWidth: '800px', margin: '0 auto' }}>
      {faqItems.map((item, index) => (
        <Accordion
          key={index}
          expanded={expanded === index}
          onChange={handleChange(index)}
          style={{ marginBottom: '10px', boxShadow: '0 2px 5px rgba(0, 0, 0, 0.1)',borderRadius: '10px' }}
        >
          <AccordionSummary
            expandIcon={<ExpandMoreIcon />}
            aria-controls={`panel${index}-content`}
            id={`panel${index}-header`}
          >
            <Typography variant="h6" component="div" style={{ fontWeight: 'normal' }}>
              {item.question}
            </Typography>
          </AccordionSummary>
          <AccordionDetails>
            <Typography>{item.answer}</Typography>
          </AccordionDetails>
        </Accordion>
      ))}
    </div>
  );
};

export default FAQSection;
