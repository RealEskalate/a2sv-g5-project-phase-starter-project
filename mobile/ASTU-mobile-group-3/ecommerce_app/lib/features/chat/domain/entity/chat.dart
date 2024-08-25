import 'package:equatable/equatable.dart';

import '../../../auth/data/model/user_model.dart';
import '../../../auth/domain/entities/user_entity.dart';

class ChatEntity extends Equatable {
  final String chatId;
  final UserEntity user1;
  final UserEntity user2;

  const ChatEntity({
    required this.chatId,
    required this.user1,
    required this.user2,
  });

  // Method to convert ChatEntity to JSON
  // Map<String, dynamic> toJson() {
  //   return {
  //     '_id': chatId,
  //     'user1': user1.toJson(),
  //     'user2': user2.toJson(),
  //   };
  // }

  factory ChatEntity.fromJson(Map<String, dynamic> json) {
    return ChatEntity(
      chatId: json['_id'],
      user1: UserModel.fromSellerJson(json['user1']),
      user2: UserModel.fromSellerJson(json['user2']),
    );
  }

  @override
  List<Object?> get props => [chatId, user1, user2];
}
