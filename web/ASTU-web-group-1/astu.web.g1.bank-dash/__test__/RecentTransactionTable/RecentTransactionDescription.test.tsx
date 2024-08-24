import React from 'react';
import { render, screen } from '@testing-library/react';
import RecentTransactionDescription from '../../src/components/RecentTransactionTable/RecentTransactionDescription';
import '@testing-library/jest-dom';

describe('RecentTransactionDescription', () => {
  it('renders the description and upward arrow for positive amount', () => {
    render(<RecentTransactionDescription amount={50} description="Income" />);

    expect(screen.getByText('Income')).toBeInTheDocument();
    expect(screen.getByTestId("upArrow").tagName.toLowerCase()).toBe("svg");
  });

  it('renders the description and downward arrow for negative amount', () => {
    render(<RecentTransactionDescription amount={-50} description="Expense" />);

    expect(screen.getByText('Expense')).toBeInTheDocument();
    expect(screen.getByTestId("downArrow").tagName.toLowerCase()).toBe('svg');
  });
});
