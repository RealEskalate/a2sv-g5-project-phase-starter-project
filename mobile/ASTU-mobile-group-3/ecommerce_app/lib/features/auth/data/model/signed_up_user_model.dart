import '../../domain/entities/signed_up_user_entity.dart';

class SignedUpUserModel extends SignedUpUserEntity {
  @override
  // ignore: overridden_fields
  final String id;
  @override
  // ignore: overridden_fields
  final String name;
  @override
  // ignore: overridden_fields
  final String email;

  const SignedUpUserModel(
      {required this.id, required this.name, required this.email})
      : super(id: id, name: name, email: email);

  @override
  List<Object?> get props => [id, name, email];

  factory SignedUpUserModel.fromJson(Map<String, dynamic> map) {
    return SignedUpUserModel(
      id: map['id'],
      name: map['name'],
      email: map['email'],
    );
  }
  SignedUpUserEntity toEntity() {
    return SignedUpUserEntity(id: id, name: name, email: email);
  }
}
