// import 'package:ecommerce_a2sv/core/network/custom_client.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:http/http.dart' as http;
// import 'package:mockito/mockito.dart';

// import '../../helpers/test_helper.mocks.dart';

// void main() {
//   late MockHttpClient mockHttpClient;
//   late CustomHttpClient customHttpClient;

//   const String authToken = 'test_token';
//   const String baseUrl = 'https://example.com/api';

//   setUp(() {
//     mockHttpClient = MockHttpClient();
//     customHttpClient = CustomHttpClient(
//       client: mockHttpClient,
//       authToken: authToken,
//       baseUrl: baseUrl,
//     );
//   });

//   group('CustomHttpClient', () {
//     const String endpoint = '/test';

//     test(
//         'get method should make a GET request with the correct headers and URL',
//         () async {
//       final uri = Uri.parse('$baseUrl$endpoint');
//       when(mockHttpClient.get(uri, headers: anyNamed('headers')))
//           .thenAnswer((_) async => http.Response('{"data": "test"}', 200));

//       final response = await customHttpClient.get(endpoint);

//       expect(response.statusCode, 200);
//       expect(response.body, '{"data": "test"}');
//       verify(mockHttpClient.get(uri, headers: {
//         'Authorization': 'Bearer $authToken',
//         'Content-Type': 'application/json',
//       })).called(1);
//     });

//     test(
//         'post method should make a POST request with the correct headers, URL, and body',
//         () async {
//       final uri = Uri.parse('$baseUrl$endpoint');
//       const body = '{"key": "value"}';
//       when(mockHttpClient.post(uri, headers: anyNamed('headers'), body: body))
//           .thenAnswer((_) async => http.Response('{"data": "test"}', 201));

//       final response = await customHttpClient.post(endpoint, body: body);

//       expect(response.statusCode, 201);
//       expect(response.body, '{"data": "test"}');
//       verify(mockHttpClient.post(uri,
//               headers: {
//                 'Authorization': 'Bearer $authToken',
//                 'Content-Type': 'application/json',
//               },
//               body: body))
//           .called(1);
//     });

//     test(
//         'put method should make a PUT request with the correct headers, URL, and body',
//         () async {
//       final uri = Uri.parse('$baseUrl$endpoint');
//       const body = '{"key": "value"}';
//       when(mockHttpClient.put(uri, headers: anyNamed('headers'), body: body))
//           .thenAnswer((_) async => http.Response('{"data": "test"}', 200));

//       final response = await customHttpClient.put(endpoint, body: body);

//       expect(response.statusCode, 200);
//       expect(response.body, '{"data": "test"}');
//       verify(mockHttpClient.put(uri,
//               headers: {
//                 'Authorization': 'Bearer $authToken',
//                 'Content-Type': 'application/json',
//               },
//               body: body))
//           .called(1);
//     });

//     test(
//         'delete method should make a DELETE request with the correct headers and URL',
//         () async {
//       final uri = Uri.parse('$baseUrl$endpoint');
//       when(mockHttpClient.delete(uri, headers: anyNamed('headers')))
//           .thenAnswer((_) async => http.Response('{"data": "test"}', 200));

//       final response = await customHttpClient.delete(endpoint);

//       expect(response.statusCode, 200);
//       expect(response.body, '{"data": "test"}');
//       verify(mockHttpClient.delete(uri, headers: {
//         'Authorization': 'Bearer $authToken',
//         'Content-Type': 'application/json',
//       })).called(1);
//     });

//     test('send method should send a request with the correct headers and URL',
//         () async {
//       final request = http.Request('POST', Uri.parse('$baseUrl$endpoint'));
//       request.body = '{"key": "value"}';

//       final streamedResponse = http.StreamedResponse(Stream.value([0, 1]), 201);
//       when(mockHttpClient.send(any)).thenAnswer((_) async => streamedResponse);

//       final response = await customHttpClient.send(request);

//       expect(response.statusCode, 201);
//       verify(mockHttpClient.send(argThat(
//         predicate<http.BaseRequest>((req) =>
//             req.url.toString() == '$baseUrl$endpoint' &&
//             req.headers['Authorization'] == 'Bearer $authToken' &&
//             req.headers['Content-Type'] == 'application/json'),
//       ))).called(1);
//     });
//   });
// }
void main(){
  
}