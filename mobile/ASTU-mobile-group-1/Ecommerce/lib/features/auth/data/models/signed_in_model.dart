import '../../domain/entities/signed_in_entity.dart';

class SignedInModel extends SignedInEntity {
  const SignedInModel({required super.accessToken});

  factory SignedInModel.fromJson(Map<String, dynamic> json) {
    return SignedInModel(
      accessToken: json['data']['access_token'],
    );
  }
}
