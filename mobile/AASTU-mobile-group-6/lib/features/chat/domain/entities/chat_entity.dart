import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:equatable/equatable.dart';

class ChatEntity extends Equatable{
  final String chatid;
  final UserModel user1;
  final UserModel user2;

  ChatEntity({required this.chatid, required this.user1, required this.user2,
  });
  
  List<Object?> get props => [chatid, user1, user2,];


  Map<String, dynamic> toJson() {
    return {
      '_id': chatid,
      'user1': user1.toJson(),
      'user2': user2.toJson(),
    };

  }

  factory ChatEntity.fromJson(Map<String, dynamic> json) {
    return ChatEntity(
      chatid: json['_id'],
      user1: UserModel.fromJson(json['user1']),
      user2: UserModel.fromJson(json['user2']),
      );
  }
  
}


