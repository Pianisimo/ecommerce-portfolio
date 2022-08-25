import React from 'react';
import './StripeButton.scss';
import StripeCheckout from "react-stripe-checkout";

const StripeButton = ({price}) => {
    const priceInCents = price * 100;
    const publishableKey = "pk_test_51La5Z5H0erQL6fArE82oDqDuiPBiyr8nxfDx0AAugJlbLYmMfjGjTR3VcHWZtPxETocX54QxWYetr3HnE6H0thvZ00YgJy26KY";

    const onToken = (token) => {
      console.log(token)
        alert('Payment Successful')
    }

    return (
        <StripeCheckout label='Pay Now'
                        name='Ecommerce Demo'
                        billingAddress
                        shippingAddress
                        image='https://svgshare.com/i/CUz.svg'
                        description={`Your total is $${price}`}
                        amount={priceInCents}
                        panelLabel='Pay Now'
                        token={onToken}
                        stripeKey={publishableKey} />
    );
};

export default StripeButton;
