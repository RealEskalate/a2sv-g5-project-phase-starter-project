import '../../../authentication/domain/entities/user_data.dart';

class Chat {
  final String id;
  final UserEntity user1;
  final UserEntity user2;

  Chat({
    required this.id,
    required this.user1,
    required this.user2,
  });
}