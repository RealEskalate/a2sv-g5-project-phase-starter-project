import 'package:ecom_app/features/chat/domain/entities/message_sender_entity.dart';

class MessageSenderModel extends MessageSenderEntity {
  final String name;
  final String email;

  MessageSenderModel({
    required this.name,
    required this.email,
  }) : super(name: name, email: email);

  // Factory constructor to create a MessageSenderModel from JSON
  factory MessageSenderModel.fromJson(Map<String, dynamic> json) {
    return MessageSenderModel(
      name: json['name'],
      email: json['email'],
    );
  }

  // Method to convert MessageSenderModel to JSON (if needed)
  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'email': email,
    };
  }
}
