import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';

import '../../../../../../core/Colors/colors.dart';
import '../../../../../../core/Icons/back_icons.dart';
import '../../../../../../core/Icons/star.dart';
import '../../../../../../core/Text_Style/text_style.dart';
import '../../../state/product_bloc/product_bloc.dart';
import '../../../state/product_bloc/product_event.dart';
import '../../../state/product_bloc/product_state.dart';
import 'delete_update_button.dart';
import 'description.dart';

class DetailPage extends StatelessWidget {
  const DetailPage({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    final Map<String, dynamic> data =
        ModalRoute.of(context)!.settings.arguments as Map<String, dynamic>;
    
    return SafeArea(
      child: Scaffold(
        appBar: PreferredSize(
          preferredSize: const Size.fromHeight(286),
          child: Stack(
            children: [
              Container(
                decoration: BoxDecoration(
                    image: DecorationImage(
                        image: NetworkImage(data['imageUrl']), fit: BoxFit.fill)),
              ),
              const BackIcons()
            ],
          ),
        ),
        body: BlocListener<ProductBloc, ProductState>(
          listener: (context, state) {
            if (state is ProductErrorState) {
              
              EasyLoading.dismiss();
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(
                      content: Text('try again'),
                    ),
                  );
                  
                  
                } else if (state is SuccessDelete) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(
                      content: Text('success'),
                    ),
                  );
                  context.read<ProductBloc>().add(const LoadAllProductEvent());
                  EasyLoading.showSuccess('success');
                  EasyLoading.dismiss();
                  Navigator.popUntil(context, ModalRoute.withName('/home'));
                }
          },
          child: SingleChildScrollView(
            child: Container(
              color: Colors.white,
              padding: const EdgeInsets.all(15),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  // short description about the product price, rating and name of brand= ====================
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Column(
                        // mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          const SizedBox(
                            height: 10,
                          ),
                          SizedBox(
                            width: MediaQuery.of(context).size.width * 0.4,
                            child: TextStyles(
                              text: data['name'], 
                              fontColor: Colors.black, 
                              fontSizes: 20,
                              fontWeight: FontWeight.bold,
                              ),
                          ),
                          
                        ],
                      ),
                      Column(
                        // mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              const Star(),
                              TextStyles(
                                text: '(4.0)',
                                fontColor: smallText,
                                fontSizes: 12,
                              )
                            ],
                          ),
                          const SizedBox(
                            height: 10,
                          ),
                          TextStyles(
                              text: '\$${data["price"]}',
                              fontColor: mainText,
                              fontSizes: 16)
                        ],
                      )
                    ],
                  ),
                  const SizedBox(
                    height: 15,
                  ),
                  // size of product if avilable
                  const SizedBox(
                    height: 10,
                  ),
                  SizedBox(
                    child: Descriptions(text: data['disc']),
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      GestureDetector(
                        onTap: () => {
                          EasyLoading.showProgress(0.3, status: 'Deleting...'),
                          context
                              .read<ProductBloc>()
                              .add(DeleteProductEvent(id: data['id'])),
                        },
                        child: DeleteUpdateButton(
                          id: data['id'],
                          text: 'DELETE',
                          bordColor: Colors.red,
                          bottonColor: Colors.white,
                        ),
                      ),
                      DeleteUpdateButton(
                        imageUrl: data['imageUrl'],
                        id: data['id'],
                        name: data['name'],
                        price: data['price'],
                        disc: data['disc'],
                        text: 'UPDATE',
                        bordColor: Colors.blue,
                        bottonColor: Colors.blue,
                      ),
                    ],
                  )
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
