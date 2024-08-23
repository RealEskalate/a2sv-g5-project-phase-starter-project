// import 'package:flutter/material.dart';
// import 'package:flutter_bloc/flutter_bloc.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';
// import 'package:product_6/features/auth/presentation/widgets/custom_outlined_button.dart';
// import 'package:product_6/features/product/presentation/bloc/product_bloc.dart';
// import 'package:product_6/features/product/presentation/pages/add_update_page.dart';

// import '../../../../helpers/test_helper.mocks.dart';

// Future<void> main() async {
//   late MockProductBloc productBloc;

//   setUp(() {
//     productBloc = MockProductBloc();
//   });

//   testWidgets('AddUpdatePage renders correctly', (WidgetTester tester) async {
//     when(() => productBloc.state).thenReturn(
//       () => ,
//     );

//     await tester.pumpWidget(
//       MaterialApp(
//         home: BlocProvider<ProductBloc>.value(
//           value: productBloc,
//           child: const AddUpdatePage(),
//         ),
//       ),
//     );

//     // Check that the AddUpdatePage widgets are present
//     expect(find.text('Add Product'), findsOneWidget);
//     expect(
//         find.byType(TextFormField), findsNWidgets(4)); // Assuming 4 TextFields
//     expect(find.byType(CustomOutlinedButton), findsOneWidget);

//     // Simulate text input
//     await tester.enterText(find.byType(TextFormField).first, 'Sample Product');
//     await tester.pump();

//     // Verify button press
//     await tester.tap(find.text('ADD'));
//     await tester.pump();
//     // Check if the Bloc event is added
//     verify(() => productBloc.add(any)).called(1);
//   });
// }
