import 'package:e_commerce_app/features/auth/data/models/user_model.dart';
import 'package:equatable/equatable.dart';

class User extends Equatable {
  final String id;
  final String email;
  final String name;

  User({
    required this.id,
    required this.email,
    required this.name,


  });
  @override
  List<Object?> get props => [id, email, name];

  static UserModel toModel(User user) {
    return UserModel(
      id: user.id,
      email: user.email,
      name: user.name,
    );
  }
}