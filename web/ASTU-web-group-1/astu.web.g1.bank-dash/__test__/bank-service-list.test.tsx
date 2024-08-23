import BankServicesList, { BankServices } from "@/components/BankServicesList/BankServicesList";
import { render, screen } from "@testing-library/react";
import '@testing-library/jest-dom';

describe("BankServicesList Component", () => {
  it("should render without crashing", () => {
    render(<BankServicesList />);
  });

  it("should render the correct number of services", async () => {
    render(<BankServicesList />);
    const serviceNames = await screen.findAllByTestId(/service-name-/);
    expect(serviceNames.length).toBe(BankServices.length);
  });

  BankServices.forEach((service, index) => {
    it(`should render correct details for service ${index + 1}`, async () => {
      render(<BankServicesList />);

      // Check for the service icon
      const icon = await screen.findByTestId(`icon-${index}`);
      expect(icon).toBeInTheDocument();

      // Check for the service name
      const serviceName = await screen.findByTestId(`service-name-${index}`);
      expect(serviceName).toHaveTextContent(service.firstcol[0]);

      // Check for the service description
      const serviceDescription = await screen.findByTestId(`service-description-${index}`);
      expect(serviceDescription).toHaveTextContent(service.firstcol[1]);

      // Check for the second column name and description (if visible)
      const secondColName = screen.queryByTestId(`second-col-name-${index}`);
      const secondColDescription = screen.queryByTestId(`second-col-description-${index}`);
      if (secondColName && secondColDescription) {
        expect(secondColName).toHaveTextContent(service.secondcol[0]);
        expect(secondColDescription).toHaveTextContent(service.secondcol[1]);
      }

      // Check for the third column name and description (if visible)
      const thirdColName = screen.queryByTestId(`third-col-name-${index}`);
      const thirdColDescription = screen.queryByTestId(`third-col-description-${index}`);
      if (thirdColName && thirdColDescription) {
        expect(thirdColName).toHaveTextContent(service.thirdcol[0]);
        expect(thirdColDescription).toHaveTextContent(service.thirdcol[1]);
      }

      // Check for the fourth column name and description (if visible)
      const fourthColName = screen.queryByTestId(`fourth-col-name-${index}`);
      const fourthColDescription = screen.queryByTestId(`fourth-col-description-${index}`);
      if (fourthColName && fourthColDescription) {
        expect(fourthColName).toHaveTextContent(service.fourthcol[0]);
        expect(fourthColDescription).toHaveTextContent(service.fourthcol[1]);
      }
    });
  });
});
