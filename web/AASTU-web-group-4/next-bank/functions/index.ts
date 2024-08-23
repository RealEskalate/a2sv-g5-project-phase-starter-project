// functions/src/index.ts
import * as functions from 'firebase-functions';
import * as admin from 'firebase-admin';
admin.initializeApp();

export const createTransactionNotification = functions.https.onCall(async (transactionData) => {
  const { senderUserId, receiverUserId, amount, transactionId } = transactionData;
  console.log('Received transaction data:', transactionData);

  const senderMessage = `You have sent $${amount} to ${receiverUserId}`;
  const receiverMessage = `You have received $${amount} from ${senderUserId}`;

  const batch = admin.firestore().batch();

  const senderNotificationRef = admin.firestore().collection('notifications').doc();
  batch.set(senderNotificationRef, {
    userId: senderUserId,
    message: senderMessage,
    timestamp: admin.firestore.FieldValue.serverTimestamp(),
    transactionId: transactionId,
    status: 'unread'
  });

  const receiverNotificationRef = admin.firestore().collection('notifications').doc();
  batch.set(receiverNotificationRef, {
    userId: receiverUserId,
    message: receiverMessage,
    timestamp: admin.firestore.FieldValue.serverTimestamp(),
    transactionId: transactionId,
    status: 'unread'
  });

  await batch.commit();
  return { success: true };
});
