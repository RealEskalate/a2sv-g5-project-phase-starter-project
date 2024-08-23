class PersonalMessageNotificationData {
  final String fullName;
  final String message;
  final String timeSent;
  final bool isRead;
  final bool isOnline;
  final int unreadMessages;

  PersonalMessageNotificationData(
      {required this.fullName,
      required this.message,
      required this.timeSent,
      required this.isRead,
      required this.isOnline,
      required this.unreadMessages});
}

List<PersonalMessageNotificationData> personalNotifications = [
  PersonalMessageNotificationData(
      fullName: 'Bereket Tesfaye',
      message: 'Hello, did you finish the task?',
      timeSent: '2:30 PM',
      isRead: false,
      isOnline: true,
    
      unreadMessages: 5),
  PersonalMessageNotificationData(
      fullName: 'Imran Mohammed',
      message: 'Dont miss the meeting',
      timeSent: '12:30 PM',
      isRead: true,
      isOnline: false,
      unreadMessages: 2),
  PersonalMessageNotificationData(
      fullName: 'Felmeta',
      message: 'Hey man long time no see',
      timeSent: '2:30 PM',
      isRead: true,
      isOnline: false,
      unreadMessages: 0),
  PersonalMessageNotificationData(
      fullName: 'Leykun ',
      message: 'wanna go out for a drink?',
      timeSent: '7:00 PM',
      isRead: true,
      isOnline: true,
      unreadMessages: 1),
  PersonalMessageNotificationData(
      fullName: 'Leul wujira',
      message: 'Did you do the assignment?',
      timeSent: '6:30 PM',
      isRead: true,
      isOnline: false,
      unreadMessages: 1),
  PersonalMessageNotificationData(
      fullName: 'My love Beth',
      message: 'I miss you, call me',
      timeSent: '1:00 AM',
      isRead: false,
      isOnline: true,
      unreadMessages: 2),
   PersonalMessageNotificationData(
      fullName: 'Felmeta',
      message: 'Hey man long time no see',
      timeSent: '2:30 PM',
      isRead: true,
      isOnline: false,
      unreadMessages: 0),
  PersonalMessageNotificationData(
      fullName: 'Leykun ',
      message: 'wanna go out for a drink sd dsdf ad fad fasd fa df?',
      timeSent: '7:00 PM',
      isRead: true,
      isOnline: true,
      unreadMessages: 1),
  PersonalMessageNotificationData(
      fullName: 'Leul wujira',
      message: 'Did you do the assignment?',
      timeSent: '6:30 PM',
      isRead: true,
      isOnline: false,
      unreadMessages: 1),
];
