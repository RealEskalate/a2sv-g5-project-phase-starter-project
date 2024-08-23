import 'package:http/http.dart' as http;
import '../../features/auth/data/data_source/auth_local_data_source.dart';

class CustomHttpClient {
  final http.Client _client;
  final AuthLocalDataSource _authLocalDataSource;

  CustomHttpClient({
    required http.Client client,
    required AuthLocalDataSource authLocalDataSource,
  })  : _client = client,
        _authLocalDataSource = authLocalDataSource;

  Future<http.Response> get(String endpoint) async {
    return _client.get(
      _parseUrl(endpoint),
      headers: await _headers(),
    );
  }

  Future<http.Response> post(String endpoint, {Object? body}) async {
    return _client.post(
      _parseUrl(endpoint),
      body: body,
      headers: await _headers(),
    );
  }

  Future<http.Response> put(String endpoint, {Object? body}) async {
    return _client.put(
      _parseUrl(endpoint),
      body: body,
      headers: await _headers(),
    );
  }

  Future<http.Response> delete(String endpoint) async {
    return _client.delete(
      _parseUrl(endpoint),
      headers: await _headers(),
    );
  }

  Future<http.StreamedResponse> send(http.BaseRequest request) async {
    request.headers.addAll(await _headers());
    return request.send();
  }

  Uri _parseUrl(String endpoint) {
    return Uri.parse(endpoint);
  }

  Future<Map<String, String>> _headers() async {
    final token = await _authLocalDataSource.getToken();
  
    return {
      'Authorization': 'Bearer $token',
      'Content-Type': 'application/json',
    };
  }
}
