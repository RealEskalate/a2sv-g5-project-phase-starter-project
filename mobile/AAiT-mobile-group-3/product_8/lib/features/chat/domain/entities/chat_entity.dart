import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user_data_entity.dart';

class ChatEntity extends Equatable {
  final String id;
  final UserDataEntity user1;
  final UserDataEntity user2;

  const ChatEntity({
    required this.id,
    required this.user1,
    required this.user2,
  });
  // add getter for id
  @override
  List<Object?> get props => [id, user1, user2];
}
