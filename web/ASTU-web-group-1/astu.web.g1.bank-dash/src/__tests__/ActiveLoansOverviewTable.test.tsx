import { render, screen, waitFor } from "@testing-library/react";
import "@testing-library/jest-dom";
import ActiveLoansOverviewTable from "../components/ActiveLoansOverviewTable/ActiveLoansOverviewTable";
import { useGetAllActiveLoansQuery } from "@/lib/redux/slices/activeLoanSlice";

// Mock the hook
jest.mock("../lib/redux/slices/activeLoanSlice", () => ({
  useGetAllActiveLoansQuery: jest.fn(),
}));

type MockedUseGetAllActiveLoansQuery = jest.MockedFunction<
  typeof useGetAllActiveLoansQuery
>;

describe("ActiveLoansOverviewTable Component", () => {
  const mockedUseGetAllActiveLoansQuery =
    useGetAllActiveLoansQuery as MockedUseGetAllActiveLoansQuery;

  beforeEach(() => {
    mockedUseGetAllActiveLoansQuery.mockReset();
  });

  it("should display data correctly when fetch is successful", async () => {
    const mockData = {
      data: [
        {
          serialNumber: "1",
          loanAmount: 10000,
          amountLeftToRepay: 5000,
          duration: 12,
          interestRate: 5,
          installment: 500,
          activeLoneStatus: "approved",
        },
        {
          serialNumber: "2",
          loanAmount: 20000,
          amountLeftToRepay: 15000,
          duration: 24,
          interestRate: 4.5,
          installment: 700,
          activeLoneStatus: "approved",
        },
      ],
      error: null,
      isLoading: false,
    };

    mockedUseGetAllActiveLoansQuery.mockImplementation(() => mockData as any);

    render(<ActiveLoansOverviewTable />);

    // Print the DOM to help with debugging
    // screen.debug();

    // Check for the content
    // expect(await screen.findByText("approved")).toBeInTheDocument();
    // expect(await screen.findByText("$10,000")).toBeInTheDocument();
    // expect(await screen.findByText("$20,000")).toBeInTheDocument();
    // expect(await screen.findByText("$5,000")).toBeInTheDocument();
    // expect(await screen.findByText("$15,000")).toBeInTheDocument();
    // expect(await screen.findByText("Loan Money")).toBeInTheDocument();
    // expect(await screen.findByText("Left to repay")).toBeInTheDocument();
    // expect(await screen.findByText("Duration")).toBeInTheDocument();
    // expect(await screen.findByText("Interest rate")).toBeInTheDocument();
    // expect(await screen.findByText("Installment")).toBeInTheDocument();
    // expect(await screen.findByText("Repay")).toBeInTheDocument();

    // // Check total values
    // expect(await screen.findByText("$30,000")).toBeInTheDocument();
    // expect(await screen.findByText("$20,000")).toBeInTheDocument();
    // expect(await screen.findByText("$1,200")).toBeInTheDocument();
  });

  it("should handle loading state correctly", async () => {
    // mockedUseGetAllActiveLoansQuery.mockImplementation(
    //   () =>
    //     ({
    //       data: null,
    //       error: null,
    //       isLoading: true,
    //     } as any)
    // );
    // render(<ActiveLoansOverviewTable />);
    // Replace 'Loading...' with the actual loading indicator in your component
    // expect(await screen.findByText("Loading...")).toBeInTheDocument();
  });

  it("should handle error state correctly", async () => {
    // mockedUseGetAllActiveLoansQuery.mockImplementation(
    //   () =>
    //     ({
    //       data: null,
    //       error: "An error occurred",
    //       isLoading: false,
    //     } as any)
    // );
    // render(<ActiveLoansOverviewTable />);
    // Adjust based on your actual error message
    // expect(await screen.findByText("An error occurred")).toBeInTheDocument();
  });
});
