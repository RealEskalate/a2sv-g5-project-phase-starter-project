import React from "react";
import {
  Page,
  Text,
  View,
  Document,
  StyleSheet,
  Image,
} from "@react-pdf/renderer";

interface ReceiptProps {
  senderUserName: string;
  receiverUserName: string;
  paymentDate: string;
  transactionId: string;
  description: string;
  amount: number;
}

const styles = StyleSheet.create({
  page: {
    padding: 30,
    fontFamily: "Helvetica",
    backgroundColor: "#f8f9fa",
  },
  logoContainer: {
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center",
    marginBottom: 20,
  },
  logo: {
    width: 60,
    height: 60,
    marginRight: 10,
  },
  bankName: {
    fontSize: 20,
    color: "#4B0082",
    fontWeight: "bold",
  },
  title: {
    fontSize: 22,
    textAlign: "center",
    color: "#4B0082",
    marginBottom: 20,
    textTransform: "uppercase",
    letterSpacing: 1,
  },
  section: {
    marginBottom: 15,
    padding: 15,
    borderRadius: 8,
    borderColor: "#4B0082",
    borderWidth: 1,
    backgroundColor: "#ffffff",
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 5,
  },
  fieldContainer: {
    flexDirection: "row",
    justifyContent: "space-between",
    marginBottom: 8,
    paddingBottom: 5,
    borderBottomColor: "#ddd",
    borderBottomWidth: 1,
  },
  fieldTitle: {
    fontSize: 14,
    fontWeight: "bold",
    color: "#4B0082",
  },
  fieldValue: {
    fontSize: 14,
    color: "#333",
  },
  footer: {
    textAlign: "center",
    fontSize: 12,
    color: "#4B0082",
    marginTop: 20,
    borderTopColor: "#ddd",
    borderTopWidth: 1,
    paddingTop: 10,
  },
});

const Receipt: React.FC<ReceiptProps> = ({
  senderUserName,
  receiverUserName,
  paymentDate,
  transactionId,
  description,
  amount,
}) => (
  <Document>
    <Page size="A4" style={styles.page}>
      <View style={styles.logoContainer}>
        <Image style={styles.logo} src="/icons/logo.png" />
        <Text style={styles.bankName}>Bank Dash</Text>
      </View>

      <Text style={styles.title}>Payment Information</Text>

      <View style={styles.section}>
        <View style={styles.fieldContainer}>
          <Text style={styles.fieldTitle}>Payer:</Text>
          <Text style={styles.fieldValue}>{senderUserName}</Text>
        </View>
        <View style={styles.fieldContainer}>
          <Text style={styles.fieldTitle}>Receiver:</Text>
          <Text style={styles.fieldValue}>{receiverUserName}</Text>
        </View>
        <View style={styles.fieldContainer}>
          <Text style={styles.fieldTitle}>Payment Date:</Text>
          <Text style={styles.fieldValue}>{paymentDate}</Text>
        </View>
        <View style={styles.fieldContainer}>
          <Text style={styles.fieldTitle}>Reference No.:</Text>
          <Text style={styles.fieldValue}>{transactionId}</Text>
        </View>
        <View style={styles.fieldContainer}>
          <Text style={styles.fieldTitle}>Reason:</Text>
          <Text style={styles.fieldValue}>{description}</Text>
        </View>
        <View style={styles.fieldContainer}>
          <Text style={styles.fieldTitle}>Transferred Amount:</Text>
          <Text style={styles.fieldValue}>${amount}</Text>
        </View>
      </View>

      <Text style={styles.footer}>
        The Bank you can always rely on.
        {"\n"}© 2024 Commercial Bank of Ethiopia. All rights reserved.
      </Text>
    </Page>
  </Document>
);

export default Receipt;
