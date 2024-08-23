import '../../domain/entities/token_entity.dart';

class TokenModel extends TokenEntity {
  @override
  // ignore: overridden_fields
  final String token;

  const TokenModel({required this.token}) : super(token: token);

  TokenEntity toEntity() {
    return TokenEntity(token: token);
  }

  factory TokenModel.fromJson(Map<String, dynamic> json) {
    return TokenModel(token: json['access_token']);
  }
}
