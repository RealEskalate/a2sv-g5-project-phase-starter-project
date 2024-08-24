import 'package:equatable/equatable.dart';

class Message extends Equatable {
  final String sender;
  final int timeStamp;
  final int unRead;
  final String body;

  const Message(
      {required this.sender,
      required this.timeStamp,
      required this.unRead,
      required this.body});

  @override
  List<Object?> get props => [sender, body, timeStamp, unRead];
}
