import 'dart:convert';
import 'package:http/http.dart' as http;
import '../../../../service_locator.dart';
import 'user_local_data_source.dart';

Future<String> fetchUserName() async {
  final storage = getIt<UserLocalDataSource>();
  String? token = await storage.getAccessToken();

  if (token == null) {
    throw Exception('No token found');
  }

  final response = await http.get(
    Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v2/users/me'),
    headers: {
      'Authorization': 'Bearer $token',
    },
  );

  if (response.statusCode == 200) {
    final data = json.decode(response.body)['data'];
    if (data['name'] != null) {
      return data['name'];
    } else {
      throw Exception('Name not found in response');
    }
  } else {
    throw Exception('Failed to load user data');
  }
}
