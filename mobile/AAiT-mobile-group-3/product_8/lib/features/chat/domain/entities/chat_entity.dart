import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user_data_entity.dart';

class ChatEntity extends Equatable {
  final String _id;
  final UserDataEntity user1;
  final UserDataEntity user2;

  const ChatEntity(
    this._id, {
    required this.user1,
    required this.user2,
  });
  // add getter for id
  String get getId => _id;
  @override
  List<Object?> get props => [_id, user1, user2];
}
