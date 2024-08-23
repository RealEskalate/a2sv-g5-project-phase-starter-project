

import 'package:equatable/equatable.dart';

import '../../../user/domain/entities/user.dart';

class ChatEntity extends Equatable{

  final String chat_id;
  final User seller_one;
  final User seller_two;
  ChatEntity({
     required this.chat_id,
    required this.seller_one,
    required this.seller_two,
  });
  @override
  List<Object?> get props => [chat_id,seller_one,seller_two];
}